import { Button, message, Menu, Dropdown } from 'antd'
import { useEffect, useState } from 'react'
import { useRecoilState, useRecoilValue, useResetRecoilState, useSetRecoilState } from 'recoil'
import { noteDBInit } from 'src/db/db'
import { NoteState, NoteType, SyncStatus } from 'src/module/note/note.model'
import {
    CloudTwoTone, SyncOutlined, EditOutlined, EyeOutlined,
    ExportOutlined, SettingTwoTone
} from '@ant-design/icons';
import noteStore from 'src/module/note/note.store'
import userStore from 'src/module/user/user.store'
import NoteBookBoard from './NoteBookBoard'
import NoteEditForm from './NoteEditForm'
import NotePage from './NotePage'
import NoteTree from './NoteTree'
import { nSQL } from '@nano-sql/core'
import dayjs from 'dayjs'
import { SYNC_NOTES } from 'src/gqls/note/mutation'
import { useMutation } from '@apollo/react-hooks'
import exportExcel from 'src/utils/excel'
const NoteIndex = () => {
    const resetCurrentNote = useResetRecoilState(noteStore.currentNote)
    const currentUser = useRecoilValue(userStore.currentUserInfo)
    const [currentNote, setCurrentNote] = useRecoilState(noteStore.currentNote)
    const setNoteTrees = useSetRecoilState(noteStore.noteTrees)
    const [noteSyncStatus, setNoteSyncStatus] = useRecoilState(noteStore.noteSyncStatus)
    const setNotes = useSetRecoilState(noteStore.notes)
    const [syncNotes] = useMutation(SYNC_NOTES);
    const [data, setData] = useState<any[]>([])
    const sync = async () => {
        if (currentUser.uid !== 0) {
            const result = await noteStore.getUnsyncedNotes(currentUser.uid)
            setNoteSyncStatus(SyncStatus.Syncing)
            try {
                const lastUpdateTime = parseInt(localStorage.getItem("note_last_update_time") || "") || 0
                await syncNote(result, lastUpdateTime)
                setNoteSyncStatus(SyncStatus.Synced)
            } catch (error) {
                setNoteSyncStatus(SyncStatus.Unsync)
            }
            if (currentNote.note_type === NoteType.Directory) {
                const notes = await noteStore.findByParentID(currentNote.id, currentUser.uid)
                setNotes(notes)
            }
            initNoteTree()
            updateCurrentNote(currentNote.id, currentNote.editable)
        }
    }
    const syncLocalNote = async () => {
        if (currentUser.uid !== 0) {
            const result = await noteStore.findAll(currentUser.uid)
            try {
                await syncNote(result, 0, true)
            } catch (error) {
                message.error("本地数据同步失败")
                return
            }
            if (currentNote.note_type === NoteType.Directory) {
                const notes = await noteStore.findByParentID(currentNote.id, currentUser.uid)
                setNotes(notes)
            }
            initNoteTree()
            updateCurrentNote(currentNote.id, currentNote.editable)
            message.success("本地数据同步成功")
        }
    }

    const syncNote = async (unsyncedNotes: any[], lastUpdateTime: number, syncLocal: boolean = false) => {
        const result = await syncNotes({
            variables: {
                "input": {
                    "lastUpdateTime": lastUpdateTime,
                    "unsyncedNotes": unsyncedNotes.map((note: any) => {
                        return {
                            id: note.id,
                            parentID: note.parent_id,
                            uid: note.uid,
                            noteType: note.note_type,
                            level: note.level,
                            title: note.title,
                            color: note.color,
                            state: note.state,
                            version: note.version,
                            createdAt: dayjs(note.created_at).unix(),
                            updatedAt: dayjs(note.updated_at).unix(),
                            content: note.content || '',
                            tags: note.tags,
                            sha1: note.sha1 || '',
                        }
                    }),
                    "syncLocal": syncLocal
                }
            }
        })
        const data = result.data.syncNotes.list
        if (data.length > 0) {
            const remoteData = data.map((v: any) => {
                return {
                    id: v.id,
                    parent_id: v.parent_id,
                    uid: v.uid,
                    note_type: v.note_type,
                    level: v.level,
                    title: v.title,
                    content: v.content,
                    tags: v.tags,
                    state: v.state,
                    version: v.version,
                    sha1: v.sha1,
                    sync_status: SyncStatus.Synced,
                    color: v.color,
                    created_at: dayjs(v.created_at * 1000).toISOString(),
                    updated_at: dayjs(v.updated_at * 1000).toISOString(),
                }
            })
            for (const remoteNote of remoteData) {
                if (remoteNote.state === NoteState.Deleted) {
                    await noteStore.deleteLocalNote(remoteNote)
                } else {
                    await nSQL("note").query('upsert', remoteNote).exec()
                }
            }
            localStorage.setItem("note_last_update_time", result.data.syncNotes.last_update_time)
        }
    }

    useEffect(() => {
        (async () => {
            await noteDBInit()
            const notes = await noteStore.findByParentID(currentNote.id, currentUser.uid)
            setNotes(notes)
            initNoteTree()
            setData(await noteStore.findAll(currentUser.uid))
            await sync()
        })()
    }, [])

    const updateCurrentNote = async (id: string, editable: boolean = false) => {
        try {
            if (id === 'root') {
                resetCurrentNote()
            } else {
                const result = await noteStore.getByID(id, currentUser.uid)
                if (result) {
                    result.editable = editable
                    setCurrentNote(result)
                }
            }
        } catch (error) {
            message.error("更新当前笔记失败：" + error)
        }
    }

    const initNoteTree = async () => {
        try {
            const result: any[] = await nSQL("note")
                .query("select")
                .where([["uid", "=", currentUser.uid], 'AND', ['state', 'IN', [NoteState.Normal]]])
                .orderBy(["updated_at DESC"])
                .exec()
            setNoteTrees(await noteStore.treeUtils(result))
        } catch (error) {
            message.error("初始化笔记树失败：" + error)
        }
    }



    const headers = [
        { title: "id", dataIndex: "id", key: "id" },
        { title: "parent_id", dataIndex: "parent_id", key: "parent_id" },
        { title: "uid", dataIndex: "uid", key: "uid" },
        { title: "note_type", dataIndex: "note_type", key: "note_type" },
        { title: "level", dataIndex: "level", key: "level" },
        { title: "title", dataIndex: "title", key: "title" },
        { title: "state", dataIndex: "state", key: "state" },
        { title: "version", dataIndex: "version", key: "version" },
        { title: "color", dataIndex: "color", key: "color" },
        { title: "content", dataIndex: "content", key: "content" },
        { title: "tags", dataIndex: "tags", key: "tags" },
        { title: "sha1", dataIndex: "sha1", key: "sha1" },
        { title: "created_at", dataIndex: "created_at", key: "created_at" },
        { title: "updated_at", dataIndex: "updated_at", key: "updated_at" },
    ]

    const menu = (
        <Menu>
            <Menu.Item>
                <Button icon={<ExportOutlined />} onClick={
                    () => { exportExcel(headers, data, currentUser.uid + "_note.xlsx") }
                } >导出全部笔记</Button>
            </Menu.Item>
            <Menu.Item>
                <Button icon={<SyncOutlined />} onClick={
                    () => { syncLocalNote() }
                } >同步本地笔记</Button>
            </Menu.Item>
        </Menu>
    )
    return (
        <div style={{
            display: "flex", flexDirection: "column", backgroundColor: "#f9f9f9"
        }} >
            <div style={{ padding: 10 }}>
                <NoteTree updateCurrentNote={updateCurrentNote} />
                <div style={{ display: 'table-cell', paddingLeft: 12 }}>
                    <Button icon={noteSyncStatus === SyncStatus.Synced ? <CloudTwoTone className="pdc-note-button-icon" /> :
                        <SyncOutlined className="pdc-note-button-icon" style={{ color: "#1890ff" }} spin={noteSyncStatus === SyncStatus.Syncing} />}
                        title={noteSyncStatus === SyncStatus.Synced ? "已同步" : "未同步"}
                        onClick={() => {
                            if (noteSyncStatus === SyncStatus.Unsync) {
                                sync()
                            }
                        }}
                    />
                    {currentNote.note_type === NoteType.File ?
                        currentNote.editable ?
                            <span><Button type="primary" icon={<EyeOutlined className="pdc-note-button-icon" />}
                                onClick={() => updateCurrentNote(currentNote.id, !currentNote.editable)} title={"预览"} /> </span> :
                            <span><Button type="primary" icon={<EditOutlined className="pdc-note-button-icon" />} onClick={() => updateCurrentNote(currentNote.id, !currentNote.editable)} title={"编辑"} />
                                {/* <a onClick={this.downloadMDFile} title={currentNote.title + ".md"} style={{ fontSize: 24, marginLeft: 4 }}><Icon type="file-markdown" theme="twoTone" /></a> */}
                                {/* <a onClick={this.downloadPDFFile} title={currentNote.title + ".pdf"} style={{ fontSize: 24, marginLeft: 4 }}><Icon type="file-pdf" theme="twoTone" /></a> */}
                            </span> : ''
                    }

                    <Dropdown overlay={menu}>
                        <SettingTwoTone className="pdc-note-button-icon" style={{ paddingLeft: 20 }} />
                    </Dropdown>
                </div>
            </div>
            {currentNote.level < 3 ? <NoteBookBoard updateCurrentNote={updateCurrentNote} initNoteTree={initNoteTree} /> : (currentNote.editable ? <NoteEditForm updateCurrentNote={updateCurrentNote} /> :
                <div style={{
                    justifyContent: "center", display: "inline-flex", marginBottom: 18,
                }}><NotePage /></div>)}
        </div >
    )
}

export default NoteIndex
