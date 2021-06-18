import { Card, Divider, message, Popover } from "antd"
import { FC, useEffect, useMemo, useState } from "react"
import { useRecoilState, useRecoilValue, useSetRecoilState } from "recoil"
import { INote, NoteState, NoteType, SyncStatus } from "src/module/note/note.model"
import noteStore from "src/module/note/note.store"
import {
    FolderTwoTone, FileTwoTone, EllipsisOutlined, PlusCircleTwoTone,
    DeleteTwoTone, CopyTwoTone, EditTwoTone, DragOutlined
} from '@ant-design/icons';
import "src/styles/card.less"
import { randomColor, shortTitle } from "src/utils/util"
import { NoteBookNewForm } from "./NoteBookNewForm"
import { nSQL } from "@nano-sql/core"
import userStore from "src/module/user/user.store"
import dayjs from "dayjs"
import modal from "antd/lib/modal"
import NoteBookEditForm from "./NoteBookEditForm"
import NoteBookMoveForm from "./NoteBookMoveForm"
import NoteMoveForm from "./NoteMoveForm"
interface INoteBookBoardProps {
    initNoteTree: () => Promise<void>
    updateCurrentNote: (id: string, editable: boolean) => Promise<void>
}

const NoteBookBoard: FC<INoteBookBoardProps> = ({ updateCurrentNote, initNoteTree }) => {
    const [notes, setNotes] = useRecoilState(noteStore.notes)
    const currentUser = useRecoilValue(userStore.currentUserInfo)
    const [currentNote, setCurrentNote] = useRecoilState(noteStore.currentNote)
    const [newBookVisible, setNewBookVisible] = useState(false)
    const [editBookVisible, setEditBookVisible] = useState(false)
    const [moveBookVisible, setMoveBookVisible] = useState(false)
    const [moveNoteVisible, setMoveNoteVisible] = useState(false)

    const setLevelOneBook = useSetRecoilState(noteStore.levelOneBooks)
    const setLevelTwoBook = useSetRecoilState(noteStore.levelTwoBooks)

    const [noteSelectIDs, setNoteSelectIDs] = useRecoilState(noteStore.noteSelectIDs)
    const [newClolor, setNewColor] = useState("")
    const [editNote, setEditNote] = useState(currentNote)
    const [moveNoteBook, setMoveNoteBook] = useState(currentNote)
    const [moveNote, setMoveNote] = useState(currentNote)
    const setNoteSyncStatus = useSetRecoilState(noteStore.noteSyncStatus)

    const showMoveBookOption = (item: INote) => {
        if (item.level !== 1 && item.note_type === NoteType.Directory) {
            return { cursor: "pointer" }
        }
        return { display: 'none' }
    }

    const showMoveNoteOption = (item: INote) => {
        if (item.level !== 1 && item.note_type === NoteType.File) {
            return { cursor: "pointer" }
        }
        return { display: 'none' }
    }

    const cards = useMemo(() => {
        return notes.map((item) =>
            <div key={item.id} className={"pdc-note-card"}
            >
                <Card
                    onClick={() => nextNoteList(item.id)}
                    bordered={false}
                    style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%', backgroundColor: '#c0c0c0' }}
                >
                    <div style={{ width: 105 }}>
                        {item.note_type === NoteType.Directory ? <FolderTwoTone
                            className="pdc-note-icon" twoToneColor={item.color} /> : <FileTwoTone className="pdc-note-icon" twoToneColor={item.color} />}
                        <div>{shortTitle(item.title, 10)}</div>
                    </div>
                </Card>
                <Popover
                    content={
                        <div >
                            <div
                                style={{ cursor: "pointer" }}
                                onClick={() => {
                                    setEditNote(item)
                                    setEditBookVisible(true)
                                }}> <EditTwoTone /> <span style={{ marginLeft: 10 }}>编辑</span></div>
                            <div onClick={() => showMoveNoteBook(item)} style={showMoveBookOption(item)}>
                                <DragOutlined style={{ color: "#1890ff" }} /><span style={{ marginLeft: 10 }}>移动文件夹</span>
                            </div>
                            <div onClick={() => showMoveNote(item)} style={showMoveNoteOption(item)}>
                                <DragOutlined style={{ color: "#1890ff" }} /><span style={{ marginLeft: 10 }}>移动笔记</span>
                            </div>
                            <div onClick={() => copyNote(item)} style={item.note_type === NoteType.File ? { cursor: "pointer" } : { display: 'none' }}>
                                <CopyTwoTone /><span style={{ marginLeft: 10 }}>创建副本</span></div>
                            <Divider style={{ margin: 4 }} />
                            <div style={{ color: 'red', cursor: "pointer" }} onClick={() => onDelete(item)
                            }><DeleteTwoTone />
                                <span style={{ marginLeft: 10 }}>删除</span></div>
                        </div>
                    }
                    placement="bottomRight"
                    trigger="hover">
                    <div className={'pdc-note-ellipse'} >
                        <EllipsisOutlined />
                    </div>
                </Popover>
            </div >)
    }, [notes])

    const onDelete = async (note: INote) => {
        let confirmText = ''
        if (note.note_type === NoteType.Directory) {
            confirmText = "确认删除文件夹[" + note.title + "]？删除后所有文件夹下的内容都无法访问"
        } else {
            confirmText = "确认删除笔记？"
        }
        modal.confirm({
            title: '确认删除?',
            okText: "确认",
            cancelText: '取消',
            content: confirmText,
            onOk: async () => {
                await noteStore.hideNote(note.id)
                await refreshNoteBoard()
                setNoteSyncStatus(SyncStatus.Unsync)
            },
        })
    }

    const addNewNotebook = async (values: any) => {
        try {
            const now = dayjs().toISOString()
            const data = await nSQL("note").query('upsert', {
                parent_id: currentNote.id,
                uid: currentUser.uid,
                note_type: NoteType.Directory,
                level: currentNote.level + 1,
                title: values.title,
                state: NoteState.Normal,
                version: 1,
                color: values.color,
                sync_status: SyncStatus.Unsync,
                created_at: now,
                updated_at: now,
            }).exec()
            const newNoteBook: any = data[0]
            // 创建LEVEL2文件夹，自动创建空的笔记
            if (newNoteBook.level === 2) {
                const result = await nSQL("note").query('upsert', {
                    parent_id: newNoteBook.id,
                    uid: currentUser.uid,
                    title: '笔记',
                    note_type: NoteType.File,
                    level: 3,
                    state: NoteState.Normal,
                    version: 1,
                    sync_status: SyncStatus.Unsync,
                    color: randomColor(),
                    created_at: now,
                    updated_at: now,
                }).exec()
                const newNote: any = result[0]
                setNoteSelectIDs([...noteSelectIDs, newNoteBook.id, newNote.id])
                await updateCurrentNote(newNote.id, true)
            } else {
                await nextNoteList(newNoteBook.id, false)
            }
        } catch (error) {
            message.error("保存文件夹出错：" + error)
        }
        await initNoteTree()
        setNoteSyncStatus(SyncStatus.Unsync)
    }

    const showMoveNoteBook = async (item: INote) => {
        setMoveNoteBook(item)
        setLevelOneBook(await noteStore.findLevelOneBooks(item.parent_id, currentUser.uid))
        setMoveBookVisible(true)
    }

    const showMoveNote = async (item: INote) => {
        setMoveNote(item)
        setLevelTwoBook(await noteStore.findLevelTwoBooks("root", currentUser.uid))
        setMoveNoteVisible(true)
    }
    const copyNote = async (item: INote) => {
        const id = await noteStore.copyNote(item, currentUser.uid)
        await refreshNoteBoard()
        setNoteSyncStatus(SyncStatus.Unsync)
        await updateCurrentNote(id, true)
    }

    const refreshNoteBoard = async () => {
        const notes = await noteStore.findByParentID(currentNote.id, currentUser.uid)
        setNotes(notes)
        initNoteTree()
    }

    const nextNoteList = async (id: string, editable: boolean = false) => {
        if (id !== "") {
            setNoteSelectIDs([...noteSelectIDs, id])
            await updateCurrentNote(id, editable)
            const notes = await noteStore.findByParentID(id, currentUser.uid)
            setNotes(notes)
        }
    }

    useEffect(() => {
        if (currentUser.uid !== 0) {
            setTimeout(async () => {
                await refreshNoteBoard()
            }, 500)
        }
    }, [currentUser])

    return (
        <div >
            <div
                key="new"
                className={"pdc-note-card"} >
                <Card
                    onClick={async () => {
                        if (currentNote.level === 2) {
                            const newNote = {
                                id: '',
                                parent_id: currentNote.id,
                                navTitle: '编辑笔记',
                                uid: currentNote.uid,
                                level: currentNote.level + 1,
                                color: randomColor(),
                                version: 1,
                                tags: [],
                                note_type: NoteType.File,
                                sync_status: SyncStatus.Unsync,
                            }
                            const id = await noteStore.insertNoteFile(newNote)
                            setNoteSyncStatus(SyncStatus.Unsync)
                            const note = await noteStore.getByID(id, currentUser.uid)
                            if (note) {
                                note.editable = true
                                setCurrentNote(note)
                                await initNoteTree()
                            }
                        } else {
                            setNewColor(randomColor())
                            setNewBookVisible(true)
                        }
                    }}
                    bordered={false}
                    style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%', backgroundColor: '#c0c0c0' }}
                >
                    <div style={{ display: 'flex', alignItems: 'center', flexDirection: 'column' }} >
                        <PlusCircleTwoTone className={"pdc-note-icon"} />
                        <div>{currentNote.level === 2 ? "新建笔记" : "新建文件夹"}</div>
                    </div>
                </Card>
            </div>
            {cards}
            <NoteBookNewForm
                visible={newBookVisible}
                newClolor={newClolor}
                addNewNotebook={addNewNotebook}
                onCancel={() => setNewBookVisible(false)} />
            <NoteBookEditForm
                visible={editBookVisible}
                note={editNote}
                refreshNoteBoard={refreshNoteBoard}
                onCancel={() => setEditBookVisible(false)} />
            <NoteBookMoveForm
                visible={moveBookVisible}
                note={moveNoteBook}
                refreshNoteBoard={refreshNoteBoard}
                onCancel={() => setMoveBookVisible(false)} />
            <NoteMoveForm
                visible={moveNoteVisible}
                note={moveNote}
                refreshNoteBoard={refreshNoteBoard}
                onCancel={() => setMoveNoteVisible(false)} />
        </div>
    )
}

export default NoteBookBoard
