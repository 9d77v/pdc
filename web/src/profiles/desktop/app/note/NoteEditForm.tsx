import { Affix, Button, Form, Input } from 'antd'
import moment from 'moment'
import * as React from 'react'
import 'src/styles/editor.less'
import { FC, useEffect, useState } from 'react'
import { useRecoilValue, useSetRecoilState } from 'recoil'
import noteStore from 'src/module/note/note.store'
import {
    UndoOutlined, RedoOutlined, CheckSquareOutlined, BoldOutlined,
    ItalicOutlined, StrikethroughOutlined, OrderedListOutlined, UnorderedListOutlined,
    CodeOutlined, LinkOutlined, TableOutlined, ClockCircleOutlined
} from '@ant-design/icons';
import NotePage from './NotePage'
import { SyncStatus } from 'src/module/note/note.model'

interface IButton {
    icon?: any
    text?: string
    type: string
    content: string
}

interface INoteEditForm {
    initNoteTree: () => Promise<void>
    updateCurrentNote: (id: string, editable: boolean) => Promise<void>
}

const NoteEditForm: FC<INoteEditForm> = ({
    initNoteTree,
    updateCurrentNote
}) => {
    const buttons: IButton[] = [
        { icon: <UndoOutlined />, type: 'undo', content: '' },
        { icon: <RedoOutlined />, type: 'redo', content: '' },
        { icon: <CheckSquareOutlined />, type: 'over', content: '\n\n- [ ] task1\n- [x] task1\n- [ ] task1\n' },
        { text: 'H1', type: 'before', content: '# ' },
        { text: 'H2', type: 'before', content: '## ' },
        { text: 'H3', type: 'before', content: '### ' },
        { text: 'H4', type: 'before', content: '#### ' },
        { text: 'H5', type: 'before', content: '##### ' },
        { text: 'H6', type: 'before', content: '###### ' },
        { icon: <BoldOutlined />, type: 'ba', content: '**' },
        { icon: <ItalicOutlined />, type: 'ba', content: '*' },
        { icon: <StrikethroughOutlined />, type: 'ba', content: '~~' },
        { icon: <OrderedListOutlined />, type: 'over', content: '\n1. 事项1\n2. 事项2\n3. 事项3' },
        { icon: <UnorderedListOutlined />, type: 'over', content: '\n- 事项1\n- 事项2\n- 事项3' },
        { icon: <CodeOutlined />, type: 'ba', content: '```\n' },
        { icon: <LinkOutlined />, type: 'before', content: '[迷之](http://www.9d77v.me "迷之")' },
        {
            icon: <TableOutlined />, type: 'over',
            content: `
        \n| Syntax      | Description | Test Text     |
        | :---        |    :----   |          :--- |
        | Header      | Title       | Here's this   |
        | Paragraph   | Text        | And more      |` },
        { icon: <ClockCircleOutlined />, type: 'clock-circle', content: '' },
    ]
    const [contentNode, setContentNode] = useState<HTMLTextAreaElement | null>(null)
    const currentNote = useRecoilValue(noteStore.currentNote)


    const [contents, setContents] = useState<string[]>([])
    const [contentIndex, setContentIndex] = useState(-1)
    const setNoteSyncStatus = useSetRecoilState(noteStore.noteSyncStatus)

    useEffect(() => {
        if (currentNote.content) {
            setContents([...contents, currentNote.content])
            setContentIndex(contentIndex + 1)
        }
    }, [])

    const [form] = Form.useForm()
    const onTitleChange = async () => {
        setTimeout(async () => {
            const title = form.getFieldValue('title')
            await noteStore.updateNoteFile(currentNote.id, title, currentNote.content || '')
            await initNoteTree()
            await updateCurrentNote(currentNote.id, true)
            setNoteSyncStatus(SyncStatus.Unsync)
        }, 300)
    }

    const onContentChange = async () => {
        setTimeout(async () => {
            const content = form.getFieldValue('content')
            setContents([...contents, content])
            setContentIndex(contents.length - 1)
            await noteStore.updateNoteFile(currentNote.id, currentNote.title || '', content)
            await updateCurrentNote(currentNote.id, true)
            setNoteSyncStatus(SyncStatus.Unsync)
        }, 300)
    }

    const onButtonClick = async (index: number) => {
        if (contentNode) {
            const data: string = form.getFieldValue('content') || ''
            const type = buttons[index].type
            let buttonContent = buttons[index].content
            let content = ''
            if (type === 'undo') {
                setContentIndex(contentIndex - 1)
                content = contents[contentIndex]
                form.setFieldsValue({
                    content
                })
                await updateCurrentNote(currentNote.id, true)
                return
            }
            if (type === 'redo') {
                setContentIndex(contentIndex + 1)
                content = contents[contentIndex]
                form.setFieldsValue({
                    content
                })
                await updateCurrentNote(currentNote.id, true)
                return
            }
            if (type === 'clock-circle') {
                buttonContent += moment().format("YYYY-MM-DD HH:mm:ss")
            }
            if (contentNode.selectionStart === contentNode.selectionEnd) {
                if (type === 'ba') {
                    content = data.substring(0, contentNode.selectionStart) +
                        buttonContent + buttonContent.split('').reverse().join('') +
                        data.substring(contentNode.selectionStart, data.length)
                } else if (type === 'before') {
                    content = data.substring(0, contentNode.selectionStart) +
                        buttonContent +
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
                    content = data.substring(0, contentNode.selectionStart) +
                        buttonContent +
                        data.substring(contentNode.selectionStart, data.length)
                } else {
                    content = data.substring(0, contentNode.selectionStart) +
                        buttonContent +
                        data.substring(contentNode.selectionEnd, data.length)
                }
            }
            setContents([...contents, content])
            setContentIndex(contents.length - 1)
            form.setFieldsValue({
                content
            })
            await updateCurrentNote(currentNote.id, true)
        }
    }

    const icons = buttons.map((v, i) => {
        return <Button key={i} icon={v.icon} onClick={onButtonClick.bind(this, i)}
            disabled={v.type === "undo" ? (contentIndex > 0 && contents.length > 1 ? false : true) :
                (v.type === "redo" ? (contentIndex + 1 < contents.length && contents.length > 1 ? false : true) : false)}
            style={{
                width: 32, height: 32, fontWeight: 500,
                justifyContent: 'center', alignItems: 'center',
                display: 'inline-flex'
            }}>{v.text}</Button>
    })

    return (
        <Affix offsetTop={64}>
            <div style={{ display: 'flex', width: '100%' }}>
                <Form
                    form={form}
                    name="noteEditForm"
                    initialValues={{ title: currentNote.title, content: currentNote.content }}
                    style={{ width: "50%", marginLeft: 10, marginRight: 10 }}
                >
                    <Form.Item
                        name="title"
                        rules={[{ required: true, message: '请输入标题!' }, {
                            max: 50, message: '标题最多50字'
                        }]}
                    >
                        <Input placeholder="标题" onChange={onTitleChange} />
                    </Form.Item>
                    {icons}
                    <Form.Item
                        name="content"
                        rules={[{ required: true, message: '请输入内容!' }, {
                            max: 10000, message: '内容最多10000字'
                        }]}
                    >
                        <textarea
                            style={{ width: "100%" }}
                            ref={node => setContentNode(node)}
                            placeholder="内容" rows={30} id='note-edit-form-content'
                            onChange={onContentChange} />
                    </Form.Item>
                </Form>
                <div style={{ flex: 1 }} >
                    <NotePage />
                </div>
            </div >
        </Affix>
    )
}

export default NoteEditForm
