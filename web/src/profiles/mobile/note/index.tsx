import { Button, Icon, Modal, NavBar, Toast } from 'antd-mobile';
import React, { useEffect, useState } from 'react';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';
import CircleButton, { ICircleButtonProps } from 'src/components/CircleButton';
import { NoteType, SyncStatus } from 'src/module/note/note.model';
import noteStore from 'src/module/note/note.store';
import {
    PlusOutlined,
    EditOutlined,
    CloudTwoTone,
    SyncOutlined
} from '@ant-design/icons';
import { noteDBInit } from 'src/db/db';
import userStore from 'src/module/user/user.store';
import { useHistory } from 'react-router-dom';
import NoteList from './NoteList';
// import { SelectKey } from 'src/constants/app';
// import { MainStore } from 'src/stores/main.store';
// import { NoteStore } from 'src/stores/note.store';
// import { CircleButton, ICircleButtonProps } from '../components/CircleButton';
// import { NoteEditForm } from './components/NoteEditForm';
// import { NoteForm } from './components/NoteForm';
// import NoteList from './components/NoteList';
// import NoteNavBar from './components/NoteNavBar';

const prompt = Modal.prompt;

const NoteIndex = () => {
    const [visible, setVisible] = useState(false)
    const [navVisible, setNavVisible] = useState(false)
    const currentUser = useRecoilValue(userStore.currentUserInfo)
    const currentNote = useRecoilValue(noteStore.currentNote)
    const [noteSyncStatus, setNoteSyncStatus] = useRecoilState(noteStore.noteSyncStatus)
    const [notes, setNotes] = useRecoilState(noteStore.notes)
    const history = useHistory()

    const sync = async () => {
        if (currentUser.uid !== "") {
            const result = await noteStore.getUnsyncedNotes(currentUser.uid)
            setNoteSyncStatus(SyncStatus.Syncing)
            await noteStore.syncNote(result)
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
        //     create_time: moment().toISOString(),
        //     update_time: moment().toISOString(),
        //     tags: [],
        //     note_type: NoteType.File,
        //     sync_status: SyncStatus.Unsync,
        //     editable: true,
        // })
    }

    const handleVisibleChange = (visible: boolean) => {
        setVisible(visible)
    }


    let data: ICircleButtonProps[] = [
        {
            right: 32,
            radius: 60,
            bottom: 0,
            display: currentNote.note_type === NoteType.Directory &&
                currentNote.level < 2 ? 'flex' : 'none',
            icon: <PlusOutlined className="pdc-note-button-icon" />,
            onClick: onNewNoteBookClick,
        },
        {
            right: 32,
            radius: 60,
            bottom: 0,
            display: currentNote.note_type === NoteType.Directory &&
                currentNote.level > 1 ? 'flex' : 'none',
            icon: <EditOutlined />,
            onClick: onNewNoteClick,
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

    const beforeNoteList = () => {

    }
    return (
        <div style={{ height: '100%', width: "100%", backgroundColor: '#fff' }}>
            <NavBar
                mode="light"
                icon={<Icon type="left"
                    onClick={() => currentNote.id === 'root' ? history.goBack() : beforeNoteList()} />}
                onLeftClick={() => beforeNoteList}
                leftContent={currentNote.id === 'root' ? '' : <span onClick={beforeNoteList}>返回</span>}
                rightContent={currentNote.note_type === NoteType.Directory ? '' : (
                    <a style={{ float: 'right', marginRight: 10, position: 'absolute', top: 10, right: 6 }}
                        onClick={() => sync()} >
                        <Button icon={noteSyncStatus === SyncStatus.Synced ? <CloudTwoTone className="pdc-note-button-icon" /> :
                            <SyncOutlined className="pdc-note-button-icon" style={{ color: "#1890ff" }} spin={noteSyncStatus === SyncStatus.Syncing} />}
                            onClick={() => {
                                if (noteSyncStatus === SyncStatus.Unsync) {
                                    sync()
                                }
                            }}
                        />
                    </ a>
                )}
            >{currentNote.navTitle}
            </NavBar>
            {currentNote.note_type === NoteType.Directory ? <NoteList /> : ""
                // currentNote.editable ? <NoteEditForm noteStore={noteStore} /> :
                //     <NoteForm noteStore={noteStore} />}
            }
            {buttons}
        </div>

    )
}

export default NoteIndex
