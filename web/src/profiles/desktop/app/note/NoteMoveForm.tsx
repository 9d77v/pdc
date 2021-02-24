import { Button, List, Modal } from 'antd'
import { FC, useState } from 'react'
import { useRecoilState, useRecoilValue } from 'recoil'
import { INote } from 'src/module/note/note.model'
import noteStore from 'src/module/note/note.store'
import userStore from 'src/module/user/user.store'

interface INoteMoveFormProps {
    note: INote,
    visible: boolean,
    onCancel: () => void,
    refreshNoteBoard: () => Promise<void>
}


const NoteMoveForm: FC<INoteMoveFormProps> = ({
    note,
    visible,
    onCancel,
    refreshNoteBoard
}) => {
    const [currentLevel, setCurrentLevel] = useState(0)
    const [listTitle, setListTitle] = useState("文件夹")
    const [levelTwoBooks, setLevelTwoBooks] = useRecoilState(noteStore.levelTwoBooks)
    const [parentID, setParentID] = useState("")
    const currentUser = useRecoilValue(userStore.currentUserInfo)

    const handleOk = async () => {
        if (parentID !== "") {
            await noteStore.moveNote(note.id, parentID)
            await refreshNoteBoard()
        }
        onCancel()
    }

    const back = async () => {
        const notes = await noteStore.findLevelTwoBooks("root", currentUser.uid)
        setLevelTwoBooks(notes)
        setCurrentLevel(0)
        setListTitle("文件夹")
    }

    const setMoveTargetBook = async (tartetNote: INote) => {
        if (tartetNote.level === 1) {
            const notes = await noteStore.findLevelTwoBooks(tartetNote.id, currentUser.uid, note.parent_id)
            setLevelTwoBooks(notes)
            setCurrentLevel(1)
            setListTitle(tartetNote.title || "")
        } else if (tartetNote.level === 2) {
            setParentID(tartetNote.id)
        }
    }

    return (
        <Modal
            visible={visible}
            title={"移动笔记"}
            okText="确认"
            onCancel={onCancel}
            onOk={handleOk}
            width={300}
            destroyOnClose={false}
            mask={true}
        >
            将 <b>{note.title}</b> 移动到另一个位置
            <div style={{ maxHeight: 300, overflowY: 'scroll' }}>
                <List
                    size="small"
                    header={<span><Button type='primary' style={{ display: currentLevel === 1 ? "" : 'none', width: 60, marginRight: 10 }}
                        onClick={() => back()}>返回</Button>{listTitle}</span>}
                    bordered={true}
                    dataSource={levelTwoBooks}
                    // tslint:disable-next-line:jsx-no-lambda
                    renderItem={(item: INote) => (<List.Item style={{ padding: 0 }}>
                        <Button style={{ width: '100%' }} onClick={() => setMoveTargetBook(item)}>{item.title}</Button>
                    </List.Item>)}
                />
            </div>
        </Modal >
    )
}


export default NoteMoveForm
