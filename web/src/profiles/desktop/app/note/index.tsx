import { Button, message } from 'antd'
import React, { useEffect } from 'react'
import { useRecoilState, useRecoilValue, useResetRecoilState, useSetRecoilState } from 'recoil'
import { noteDBInit } from 'src/db/db'
import { NoteState, NoteType, SyncStatus } from 'src/module/note/note.model'
import {
    CloudTwoTone, SyncOutlined, EditOutlined, EyeOutlined
} from '@ant-design/icons';
import noteStore from 'src/module/note/note.store'
import userStore from 'src/module/user/user.store'
import NoteBookBoard from './NoteBookBoard'
import NoteEditForm from './NoteEditForm'
import NotePage from './NotePage'
import NoteTree from './NoteTree'
import { nSQL } from '@nano-sql/core'
const NoteIndex = () => {
    const resetCurrentNote = useResetRecoilState(noteStore.currentNote)
    const currentUser = useRecoilValue(userStore.currentUserInfo)
    const [currentNote, setCurrentNote] = useRecoilState(noteStore.currentNote)
    const setNoteTrees = useSetRecoilState(noteStore.noteTrees)
    const [noteSyncStatus, setNoteSyncStatus] = useRecoilState(noteStore.noteSyncStatus)
    const setNotes = useSetRecoilState(noteStore.notes)

    const sync = async () => {
        if (currentUser.uid !== "") {
            const result = await noteStore.getUnsyncedNotes(currentUser.uid)
            setNoteSyncStatus(SyncStatus.Syncing)
            await noteStore.syncNote(result)
            initNoteTree()
            if (currentNote.note_type === NoteType.Directory) {
                const notes = await noteStore.findByParentID(currentNote.id, currentUser.uid)
                setNotes(notes)
            }
            setNoteSyncStatus(SyncStatus.Synced)
        }
    }
    useEffect(() => {
        noteDBInit()
        setTimeout(sync, 1000)
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

    return (
        <div style={{ display: "flex", flexDirection: "column" }} >
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
                </div>
            </div>
            {currentNote.level < 3 ? <NoteBookBoard updateCurrentNote={updateCurrentNote} initNoteTree={initNoteTree} /> : (currentNote.editable ? <NoteEditForm updateCurrentNote={updateCurrentNote} initNoteTree={initNoteTree} /> : <NotePage />)}
        </div >
    )
}

export default NoteIndex
