import { Button, List, Modal } from 'antd'
import { FC, useState } from 'react';
import { useRecoilValue } from 'recoil';
import { INote } from 'src/module/note/note.model';
import noteStore from 'src/module/note/note.store';

interface INoteBookMoveFormProps {
    note: INote,
    visible: boolean,
    onCancel: () => void,
    refreshNoteBoard: () => Promise<void>
}

const NoteBookMoveForm: FC<INoteBookMoveFormProps> = ({
    note,
    visible,
    onCancel,
    refreshNoteBoard
}) => {
    const levelOneBooks = useRecoilValue(noteStore.levelOneBooks)
    const [parentID, setParentID] = useState("")

    const handleOk = async () => {
        if (parentID !== "") {
            await noteStore.moveNote(note.id, parentID)
            await refreshNoteBoard()
        }
        onCancel()
    }

    return (
        <Modal
            visible={visible}
            title={"移动文件夹"}
            okText="确认"
            cancelText="取消"
            onCancel={onCancel}
            onOk={handleOk}
            width={450}
            destroyOnClose={false}
            getContainer={false}
            mask={true}
        >
            <span style={{ paddingBottom: 32 }}>将 <b>{note.title}</b> 移动到另一个位置</span>
            <div style={{ maxHeight: 300, overflowY: 'scroll' }}>
                <List
                    size="small"
                    header={<div>文件夹</div>}
                    bordered={true}
                    dataSource={levelOneBooks}
                    // tslint:disable-next-line:jsx-no-lambda
                    renderItem={(item: INote) => (<List.Item style={{ padding: 0 }}>
                        <Button style={{ width: '100%' }} onClick={() => setParentID(item.id)}>{item.title}</Button>
                    </List.Item>)}
                />
            </div>
        </Modal >
    )
}

export default NoteBookMoveForm
