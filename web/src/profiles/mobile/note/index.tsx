import { Icon, Modal, NavBar, Toast } from 'antd-mobile';
import React, { useEffect } from 'react';
import { useRecoilState, useRecoilValue, useResetRecoilState, useSetRecoilState } from 'recoil';
import CircleButton, { ICircleButtonProps } from 'src/components/CircleButton';
import { NoteState, NoteType, SyncStatus } from 'src/module/note/note.model';
import noteStore from 'src/module/note/note.store';
import {
    PlusOutlined,
    EditOutlined,
    CloudTwoTone,
    SyncOutlined,
    EyeOutlined,
} from '@ant-design/icons'
import { noteDBInit } from 'src/db/db'
import userStore from 'src/module/user/user.store'
import { useHistory } from 'react-router-dom'
import NoteList from './NoteList'
import { Button, message } from 'antd'
import NotePage from 'src/profiles/desktop/app/note/NotePage'
import NoteEditForm from './NoteEditForm';
import { nSQL } from '@nano-sql/core';
import dayjs from 'dayjs';
import { SYNC_NOTES } from 'src/gqls/note/mutation';
import { useMutation } from '@apollo/react-hooks';

const prompt = Modal.prompt;

const NoteIndex = () => {
    const currentUser = useRecoilValue(userStore.currentUserInfo)
    const resetCurrentNote = useResetRecoilState(noteStore.currentNote)
    const [currentNote, setCurrentNote] = useRecoilState(noteStore.currentNote)
    const [noteSyncStatus, setNoteSyncStatus] = useRecoilState(noteStore.noteSyncStatus)
    const setNotes = useSetRecoilState(noteStore.notes)
    const history = useHistory()

    const [syncNotes] = useMutation(SYNC_NOTES);

    const sync = async () => {
        if (currentUser.uid !== "") {
            const result = await noteStore.getUnsyncedNotes(currentUser.uid)
            setNoteSyncStatus(SyncStatus.Syncing)
            await syncNote(result)
            initNoteTree()
            if (currentNote.note_type === NoteType.Directory) {
                const notes = await noteStore.findByParentID(currentNote.id, currentUser.uid)
                setNotes(notes)
            }
            setNoteSyncStatus(SyncStatus.Synced)
        }
    }

    const syncNote = async (unsyncedNotes: any[]) => {
        const lastUpdateTime = parseInt(localStorage.getItem("note_last_update_time") || "") || 0
        const result = await syncNotes({
            variables: {
                "input": {
                    "lastUpdateTime": lastUpdateTime === null ? 0 : lastUpdateTime,
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
                    })
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
            await sync()
        })()
    }, [])

    const updateCurrentNote = async (id: string, editable: boolean = false, navTitle?: string) => {
        try {
            if (id === 'root') {
                resetCurrentNote()
            } else {
                const result = await noteStore.getByID(id, currentUser.uid)
                if (result) {
                    result.editable = editable
                    result.navTitle = navTitle || ""
                    setCurrentNote(result)
                }
            }
        } catch (error) {
            message.error("更新当前笔记失败：" + error)
        }
    }

    const onNewNoteBookClick = () => {
        prompt(
            '新建笔记本',
            '',
            async (value) => {
                if (value.length > 20) {
                    Toast.info('标题长度不能超过20', 1);
                } else {
                    // await this.props.noteStore.addNewNotebook(value, '')
                }
            },
            'default',
            '',
            ['请输入名称'], /Android/i.test(navigator.userAgent) ? 'android' : 'ios'
        )
    }

    const onNewNoteClick = () => {
        // this.props.noteStore.navNotes.push({
        //     id: '',
        //     parent_id: this.props.noteStore.currentNote.id,
        //     navTitle: '编辑笔记',
        //     uid: this.props.mainStore.currentUser.uid,
        //     level: this.props.noteStore.currentNote.level + 1,
        //     version: 1,
        //     create_time: dayjs().toISOString(),
        //     update_time: dayjs().toISOString(),
        //     tags: [],
        //     note_type: NoteType.File,
        //     sync_status: SyncStatus.Unsync,
        //     editable: true,
        // })
    }

    let data: ICircleButtonProps[] = [
        // {
        //     right: 32,
        //     radius: 60,
        //     bottom: 0,
        //     display: currentNote.note_type === NoteType.Directory &&
        //         currentNote.level < 2 ? 'flex' : 'none',
        //     icon: <PlusOutlined className="pdc-note-button-icon" />,
        //     onClick: onNewNoteBookClick,
        // },
        // {
        //     right: 32,
        //     radius: 60,
        //     bottom: 0,
        //     display: (currentNote.note_type === NoteType.Directory && currentNote.level === 2) ||
        //         (currentNote.note_type === NoteType.File && !currentNote.editable) ? 'flex' : 'none',
        //     icon: <EditOutlined />,
        //     onClick: () => {
        //         if (currentNote.note_type === NoteType.Directory && currentNote.level === 2) {

        //         } else if (currentNote.note_type === NoteType.File && !currentNote.editable) {
        //             updateCurrentNote(currentNote.id, true, currentNote.navTitle)
        //         }
        //     },
        // },
        {
            right: 32,
            radius: 60,
            bottom: 0,
            display: (currentNote.note_type === NoteType.File && currentNote.editable) ? 'flex' : 'none',
            icon: <EyeOutlined />,
            onClick: () => {
                updateCurrentNote(currentNote.id, false, currentNote.navTitle)
            },
        }
    ]
    data = data.filter(v => v.display === 'flex')
    for (let i = 0; i < data.length; i++) {
        data[i].bottom = 66 + 76 * i
    }
    const buttons = data.map((value: ICircleButtonProps, index: number) => {
        return (
            <CircleButton
                key={index}
                right={value.right}
                bottom={value.bottom}
                radius={value.radius}
                display={value.display}
                icon={value.icon}
                onClick={value.onClick} />
        )
    })

    const beforeNoteList = async () => {
        const notes = await noteStore.findByParentID(currentNote.parent_id, currentUser.uid)
        setNotes(notes)
        const note = await noteStore.getByID(currentNote.parent_id, currentUser.uid)
        await updateCurrentNote(note?.id || 'root', false, note?.title || '记事本')
    }

    return (
        <div style={{
            height: '100%', width: "100%", backgroundColor: '#fff', display: "flex",
            flexDirection: "column", overflowY: "scroll"
        }}>
            <NavBar
                mode="light"
                icon={<Icon type="left"
                    onClick={() => currentNote.id === 'root' ? history.goBack() : beforeNoteList()} />}
                onLeftClick={() => beforeNoteList()}
                leftContent={currentNote.id === 'root' ? '' : <span onClick={beforeNoteList}>{currentNote.navTitle}</span>}
                rightContent={
                    // {currentNote.note_type === NoteType.File ?
                    //     currentNote.editable ?
                    //         <span style={{ marginRight: 10 }}><Button type="primary" icon={<EyeOutlined className="pdc-note-button-icon" />}
                    //             onClick={() => updateCurrentNote(currentNote.id, currentNote.navTitle, !currentNote.editable)} /> </span> :
                    //         <span><Button type="primary" icon={<EditOutlined className="pdc-note-button-icon" />} onClick={() => updateCurrentNote(currentNote.id, currentNote.navTitle, !currentNote.editable)} />
                    //             {/* <a onClick={this.downloadMDFile} title={currentNote.title + ".md"} style={{ fontSize: 24, marginLeft: 4 }}><Icon type="file-markdown" theme="twoTone" /></a> */}
                    //             {/* <a onClick={this.downloadPDFFile} title={currentNote.title + ".pdf"} style={{ fontSize: 24, marginLeft: 4 }}><Icon type="file-pdf" theme="twoTone" /></a> */}
                    //         </span> : ''
                    // }
                    <Button icon={noteSyncStatus === SyncStatus.Synced ? <CloudTwoTone className="pdc-note-button-icon" /> :
                        <SyncOutlined className="pdc-note-button-icon" style={{ color: "#1890ff" }} spin={noteSyncStatus === SyncStatus.Syncing} />}
                        onClick={() => {
                            if (noteSyncStatus === SyncStatus.Unsync) {
                                sync()
                            }
                        }}
                    />
                }
            >{currentNote.id === 'root' ? "记事本" : ""}
            </NavBar>
            {currentNote.note_type === NoteType.Directory ? <NoteList updateCurrentNote={updateCurrentNote} /> : (currentNote.editable ? <NoteEditForm updateCurrentNote={updateCurrentNote} /> : <NotePage />)}
            {buttons}
        </div>

    )
}

export default NoteIndex
function initNoteTree() {
    throw new Error('Function not implemented.');
}

