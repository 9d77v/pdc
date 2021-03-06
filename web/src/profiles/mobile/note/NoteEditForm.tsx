import { Button, Form, Input } from 'antd'
import dayjs from "dayjs";
import * as React from 'react'
import 'src/styles/editor.less'
import { FC, useState } from 'react'
import { useRecoilValue, useSetRecoilState } from 'recoil'
import noteStore from 'src/module/note/note.store'
import {
    CheckSquareOutlined, BoldOutlined,
    ItalicOutlined, StrikethroughOutlined, OrderedListOutlined, UnorderedListOutlined,
    CodeOutlined, LinkOutlined, TableOutlined, ClockCircleOutlined
} from '@ant-design/icons';
import { SyncStatus } from 'src/module/note/note.model'

interface IButton {
    icon?: any
    text?: string
    type: string
    content: string
}

interface INoteEditForm {
    updateCurrentNote: (id: string, editable: boolean, navTitle: string) => Promise<void>
}

const NoteEditForm: FC<INoteEditForm> = ({
    updateCurrentNote
}) => {
    const buttons: IButton[] = [
        { icon: <CheckSquareOutlined />, type: 'over', content: '- [ ] 任务1\n- [x] 任务2\n- [ ] 任务3\n' },
        { text: 'H', type: 'before', content: '#' },
        { icon: <BoldOutlined />, type: 'ba', content: '**' },
        { icon: <ItalicOutlined />, type: 'ba', content: '*' },
        { icon: <StrikethroughOutlined />, type: 'ba', content: '~~' },
        { icon: <OrderedListOutlined />, type: 'over', content: '\n1. 事项1\n2. 事项2\n3. 事项3' },
        { icon: <UnorderedListOutlined />, type: 'over', content: '\n- 事项1\n- 事项2\n- 事项3' },
        { icon: <CodeOutlined />, type: 'ba', content: '```\n' },
        { icon: <LinkOutlined />, type: 'over', content: '[个人数据中心](https://github.com/9d77v/pdc "Github地址")' },
        {
            icon: <TableOutlined />, type: 'over',
            content: `
|序号| 问 | 答 |
| :---| :---| :--- |
| 1  | Who are you? |I'm your friend.|
| 2  | Where are you?|I'm here.|` },

        { icon: <ClockCircleOutlined />, type: 'clock-circle', content: '' },
    ]
    const [contentNode, setContentNode] = useState<HTMLTextAreaElement | null>(null)
    const currentNote = useRecoilValue(noteStore.currentNote)

    const setNoteSyncStatus = useSetRecoilState(noteStore.noteSyncStatus)

    const [form] = Form.useForm()
    const onTitleChange = async () => {
        setTimeout(async () => {
            const title = form.getFieldValue('title')
            await noteStore.updateNoteFile(currentNote.id, title, currentNote.content || '')
            await updateCurrentNote(currentNote.id, true, title)
            setNoteSyncStatus(SyncStatus.Unsync)
        }, 300)
    }

    const onContentChange = async () => {
        setTimeout(async () => {
            const content = form.getFieldValue('content')
            await noteStore.updateNoteFile(currentNote.id, currentNote.title || '', content)
            await updateCurrentNote(currentNote.id, true, currentNote.title || '')
            setNoteSyncStatus(SyncStatus.Unsync)
        }, 300)
    }

    const onButtonClick = async (index: number) => {
        if (contentNode) {
            const data: string = form.getFieldValue('content') || ''
            const type = buttons[index].type
            let buttonContent = buttons[index].content
            let content = ''
            if (type === 'clock-circle') {
                buttonContent += dayjs().format("YYYY-MM-DD HH:mm:ss")
            }
            if (contentNode.selectionStart === contentNode.selectionEnd) {
                if (type === 'ba') {
                    content = data.substring(0, contentNode.selectionStart) +
                        buttonContent + buttonContent.split('').reverse().join('') +
                        data.substring(contentNode.selectionStart, data.length)
                } else if (type === 'before') {
                    content = data.substring(0, contentNode.selectionStart) +
                        buttonContent + " 标题" +
                        data.substring(contentNode.selectionStart, data.length)
                } else {
                    content = data.substring(0, contentNode.selectionStart) + buttonContent +
                        data.substring(contentNode.selectionStart, data.length)
                }
            } else {
                if (type === 'ba') {
                    content = data.substring(0, contentNode.selectionStart) +
                        buttonContent +
                        data.substring(contentNode.selectionStart, contentNode.selectionEnd) +
                        buttonContent.split('').reverse().join('') +
                        data.substring(contentNode.selectionEnd, data.length)
                } else if (type === 'before') {
                    content = data.substring(0, contentNode.selectionStart - 1) +
                        buttonContent + " " +
                        data.substring(contentNode.selectionStart, data.length)
                } else {
                    content = data.substring(0, contentNode.selectionStart) +
                        buttonContent +
                        data.substring(contentNode.selectionEnd, data.length)
                }
            }
            form.setFieldsValue({
                content
            })
            await onContentChange()
        }
    }

    const icons = buttons.map((v, i) => {
        return <Button key={i} icon={v.icon} onClick={onButtonClick.bind(this, i)}
            style={{
                width: 32, height: 32, fontWeight: 500,
                justifyContent: 'center', alignItems: 'center',
                display: 'flex', float: 'left'
            }}>{v.text}</Button>
    })

    return (
        <Form
            form={form}
            name="noteEditForm"
            initialValues={{ title: currentNote.title, content: currentNote.content }}
            style={{ padding: 5 }}
        >
            <Form.Item
                name="title"
                rules={[{ required: true, message: '请输入标题!' }, {
                    max: 50, message: '标题最多50字'
                }]}
            >
                <Input placeholder="标题" onChange={onTitleChange} />
            </Form.Item>
            <div style={{ display: "flex", flexWrap: "wrap" }}>
                {icons}
            </div>
            <Form.Item
                name="content"
                rules={[{ required: true, message: '请输入内容!' }, {
                    max: 10000, message: '内容最多10000字'
                }]}
            >
                <textarea
                    style={{ width: "100%", overflow: "auto" }}
                    ref={node => setContentNode(node)}
                    placeholder="内容" rows={30} id='note-edit-form-content'
                    onChange={onContentChange} />
            </Form.Item>
        </Form>
    )
}

export default NoteEditForm
