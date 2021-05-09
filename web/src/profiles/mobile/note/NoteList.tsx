import { List, Modal, SwipeAction, Toast } from 'antd-mobile';
import { FC } from 'react';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';
import { LongPressAction } from 'src/components';
import { INote, NoteType, SyncStatus } from 'src/module/note/note.model';
import noteStore from 'src/module/note/note.store';
import userStore from 'src/module/user/user.store';
import {
    FolderTwoTone, FileTwoTone
} from '@ant-design/icons';
const Item = List.Item;
const alert = Modal.alert
const prompt = Modal.prompt
const operation = Modal.operation

interface INoteListProps {
    updateCurrentNote: (id: string, editable?: boolean, navTitle?: string,) => Promise<void>
}
const NoteList: FC<INoteListProps> = ({
    updateCurrentNote
}) => {
    const [notes, setNotes] = useRecoilState(noteStore.notes)
    const currentUser = useRecoilValue(userStore.currentUserInfo)
    const setNoteSyncStatus = useSetRecoilState(noteStore.noteSyncStatus)
    const currentNote = useRecoilValue(noteStore.currentNote)

    const nextNoteList = async (item: INote, editable: boolean = false) => {
        if (item.id !== "") {
            await updateCurrentNote(item.id, editable, item.title)
            const notes = await noteStore.findByParentID(item.id, currentUser.uid)
            setNotes(notes)
        }
    }

    const onLongPress = (note: INote) => {
        operation([
            { text: '重命名', onPress: () => onRename(note) },
            { text: '删除', onPress: () => onDelete(note) },
        ])
    }

    const onRename = (note: INote) => {
        prompt(
            '新标题',
            '',
            async (title) => {
                if (title.length > 20) {
                    Toast.info('标题长度不能超过20', 1);
                } else {
                    await noteStore.updateNoteBrief(note.id, title, note.color || "")
                    const notes = await noteStore.findByParentID(currentNote.id, currentUser.uid)
                    setNotes(notes)
                    setNoteSyncStatus(SyncStatus.Unsync)
                }
            },
            'default',
            note.title,
            ['请输入标题'], /Android/i.test(navigator.userAgent) ? 'android' : 'ios'
        )
    }

    const onDelete = (note: INote) => {
        let confirmText = ""
        if (note.note_type === NoteType.Directory) {
            confirmText = '确认删除笔记本[' + note.title + ']？删除后笔记本下所有的内容都无法访问'
        } else {
            confirmText = "确认删除笔记？"
        }

        alert('确认删除', confirmText, [
            {
                text: '确认', onPress: async () => {
                    await noteStore.hideNote(note.id)
                    const notes = await noteStore.findByParentID(currentNote.id, currentUser.uid)
                    setNotes(notes)
                    setNoteSyncStatus(SyncStatus.Unsync)
                }
            },
            { text: '取消' }
        ],
            /Android/i.test(navigator.userAgent) ? 'android' : 'ios'
        )
    }

    let listItems: any[] = []
    if (notes) {
        listItems = notes.map((item: INote, i: number) => {
            const listItem = <Item
                style={{ height: 60 }}
                onClick={() => nextNoteList(item)}
            >
                {item.note_type === NoteType.Directory ? <FolderTwoTone
                    className="pdc-note-icon" twoToneColor={item.color} /> : <FileTwoTone className="pdc-note-icon" twoToneColor={item.color} />}
                <span style={{ marginLeft: 10 }}> {item.title === undefined ? '' : item.title}</span>
            </Item>
            if (/Android/i.test(navigator.userAgent)) {
                return (
                    <LongPressAction key={item.id}
                        timeout={500}
                        onLongPress={onLongPress.bind(this, item)} item={listItem}

                    />
                )
            }
            return (
                <SwipeAction key={item.id}
                    style={{ backgroundColor: 'gray' }}
                    autoClose={true}
                    right={[
                        {
                            text: '重命名',
                            onPress: () => onRename(item),
                            style: { backgroundColor: '#ddd', color: 'white', width: 60 },
                        },
                        {
                            text: '删除',
                            onPress: () => onDelete(item),
                            style: { backgroundColor: '#F4333C', color: 'white', width: 60 },
                        },
                    ]}
                >
                    {listItem}
                </SwipeAction>
            )
        }
        )
    }
    return (
        <List style={{ overflowY: "scroll" }}>
            {listItems}
        </List>
    )
}


export default NoteList
