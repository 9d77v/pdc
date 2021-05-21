import { INote, INoteTree, NoteState, NoteType, SyncStatus } from "./note.model"
import {
    atom,
} from 'recoil';
import { nSQL } from "@nano-sql/core";
import { getUID, randomColor, Sha1 } from "src/utils/util";
import React from "react";
import {
    FolderTwoTone, FileTwoTone,
} from '@ant-design/icons';
import dayjs from "dayjs";
import { diaryRegex } from "src/consts/regex";
require('dayjs/locale/zh-cn')
dayjs.locale('zh-cn')
class NoteStore {
    currentNote = atom<INote>({
        key: 'currentNote',
        default: {
            id: 'root',
            uid: getUID(),
            parent_id: '',
            navTitle: '记事本',
            level: 0,
            version: 1,
            sync_status: SyncStatus.Synced,
            note_type: NoteType.Directory,
        },
    })
    noteSyncStatus = atom<SyncStatus>({
        key: 'noteSynced',
        default: SyncStatus.Unsync
    })
    noteTrees = atom<INoteTree[]>({
        key: 'noteTrees',
        default: []
    })
    noteSelectIDs = atom<(string | number)[]>({
        key: 'noteSelectIDs',
        default: [],
    })
    notes = atom<INote[]>({
        key: 'notes',
        default: []
    })
    levelOneBooks = atom<INote[]>({
        key: 'levelOneBooks',
        default: []
    })
    levelTwoBooks = atom<INote[]>({
        key: 'levelTwoBooks',
        default: []
    })

    getByID = async (id: string, uid: string) => {
        const result: any[] = await nSQL("note")
            .query("select")
            .where([["id", "=", id], 'AND', ["uid", "=", uid]])
            .exec()
        if (result.length < 1) {
            return undefined
        }
        return result[0]
    }

    findAll = async (uid: string) => {
        const result: any[] = await nSQL("note")
            .query("select")
            .where([["uid", "=", uid]])
            .orderBy(["updated_at DESC"])
            .exec()
        return result
    }

    findByParentID = async (id: string, uid: string) => {
        const result: any[] = await nSQL("note")
            .query("select")
            .where([["parent_id", "=", id], 'AND', ["uid", "=", uid], 'AND', ['state', 'IN', [NoteState.Normal]]])
            .orderBy(["updated_at DESC"])
            .exec()
        return result
    }

    findLevelOneBooks = async (noteID: string, uid: string) => {
        const result: any[] = await nSQL("note")
            .query("select")
            .where([["level", "=", 1], 'AND', ["uid", "=", uid], 'AND',
            ['state', 'IN', [NoteState.Normal]], 'AND', ['id', '!=', noteID]])
            .orderBy(["updated_at DESC"])
            .exec()
        return result
    }

    findLevelTwoBooks = async (noteID: string, uid: string, omitNoteID: string = "") => {
        const result: any[] = await nSQL("note")
            .query("select")
            .where([["parent_id", "=", noteID],
                'AND', ["uid", "=", uid],
                'AND', ['state', 'IN', [NoteState.Normal]],
                'AND', ['id', '!=', omitNoteID]])
            .orderBy(["updated_at DESC"])
            .exec()
        return result
    }

    getUnsyncedNotes = async (uid: string) => {
        const result: any[] = await nSQL("note")
            .query("select")
            .where([["uid", "=", uid], 'AND', ['sync_status', 'IN', [SyncStatus.Unsync]]])
            .exec()
        return result
    }
    insertNoteFile = async (note: INote): Promise<string> => {
        const now = dayjs().toISOString()
        let data: any[] = await nSQL("note").query('upsert', {
            parent_id: note.parent_id,
            uid: note.uid,
            title: "笔记",
            note_type: NoteType.File,
            level: note.level,
            state: NoteState.Normal,
            version: 1,
            sync_status: SyncStatus.Unsync,
            color: note.color,
            created_at: now,
            updated_at: now,
        }).exec()
        return data[0].id
    }

    updateNoteFile = async (id: string, title: string, content: string) => {
        let data: any[] = await nSQL("note").query('upsert', {
            id: id,
            title: title,
            content: content,
            sha1: await Sha1(content),
            sync_status: SyncStatus.Unsync,
            updated_at: dayjs().toISOString(),
        }).exec()
        return data[0]
    }

    updateNoteBrief = async (id: string, title: string, color: string) => {
        await nSQL("note").query('upsert', {
            id: id,
            title: title,
            sync_status: SyncStatus.Unsync,
            color: color,
            updated_at: dayjs().toISOString(),
        }).exec()
    }

    moveNote = async (id: string, parentID: string) => {
        await nSQL("note").query('upsert', {
            id: id,
            parent_id: parentID,
            sync_status: SyncStatus.Unsync,
            updated_at: dayjs().toISOString(),
        }).exec()
    }

    //新副本标题，
    //1.默认最前方添加‘副本’二字
    //2.存在数字时，返回数字加1
    //3.为日期时返回下一日起
    newCopyTitle = (title: string) => {
        let newTitle = "副本 " + title
        const reg = /\d+(\.\d+)?/g
        const arr = title?.match(reg)
        if (arr?.length === 1) {
            newTitle = title?.replace(reg, (parseInt(arr[0]) + 1).toString()) || "副本 " + title
        } else if (arr?.length === 2 && diaryRegex.test(title)) {
            const year = dayjs().year()
            let month = arr[0]
            let day = arr[1]
            if (month.length === 1) {
                month = parseInt('0' + month).toString()
            }
            if (day.length === 1) {
                day = parseInt('0' + day).toString()
            }
            newTitle = dayjs(year + '-' + month + '-' + day).add(1, 'day').format('MM月DD日 dddd')
        }
        return newTitle
    }

    copyNote = async (note: INote, uid: string) => {
        const now = dayjs().toISOString()

        let data: any[] = await nSQL("note").query('upsert', {
            parent_id: note.parent_id,
            uid: uid,
            note_type: NoteType.File,
            level: note.level,
            title: this.newCopyTitle(note.title || ''),
            content: note.content,
            sha1: note.sha1,
            tags: note.tags,
            state: note.state,
            version: 1,
            sync_status: SyncStatus.Unsync,
            color: randomColor(),
            created_at: now,
            updated_at: now,
        }).exec()
        return data[0].id
    }

    hideNote = async (id: string) => {
        await nSQL("note").query('upsert', {
            id: id,
            state: NoteState.Deleted,
            sync_status: SyncStatus.Unsync,
            updated_at: dayjs().toISOString(),
        }).exec()
    }

    deleteLocalNote = async (note: any) => {
        nSQL("note").query("delete").where(["id", "=", note.id]).exec()
        if (note.note_type === NoteType.Directory) {
            const notes = await nSQL("note")
                .query("select")
                .where(["parent_id", "=", note.id])
                .exec()
            notes.forEach((item: any) => {
                this.deleteLocalNote(item)
            })
        }
    }

    private sonsTree = (obj: INoteTree, data: INote[]): INoteTree => {
        const children: INoteTree[] = [];
        data.forEach((d: INote) => {
            if (d.parent_id === obj.value) {
                const note = this.getLabel(d)
                const o = this.sonsTree(note, data);
                children.push(o);
            }
        })
        if (children.length > 0) {
            obj.children = children;
        }
        return obj;
    }

    treeUtils = (data: INote[]) => {
        const ptree: INoteTree[] = [];
        data.forEach((d: INote) => {
            if (d.parent_id === 'root') {
                const note = this.getLabel(d)
                const o = this.sonsTree(note, data);
                ptree.push(o);
            }
        });
        return ptree
    }

    getLabel = (d: INote): INoteTree => {
        const iconType = d.note_type === NoteType.Directory ? FolderTwoTone : FileTwoTone
        return {
            value: d.id,
            title: d.title || '',
            noteType: d.note_type,
            color: d.color || "",
            label: React.createElement('span',
                { display: 'flex', "align-items": 'center' },
                React.createElement(iconType, {
                    twoToneColor: d.color || '',
                    className: "pdc-note-cascader-icon",
                    key: "icon-" + d.id
                }),
                React.createElement('span', {
                    key: 'span-' + d.id,
                }, " " + d.title || ''),
            )
        }
    }
}

const noteStore = new NoteStore()
export default noteStore
