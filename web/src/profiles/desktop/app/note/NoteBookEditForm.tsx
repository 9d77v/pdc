import { Form, Input, Modal } from 'antd'
import * as React from 'react'
import { FC, useEffect, useState } from 'react';
import { SketchPicker } from 'react-color'
import { INote } from 'src/module/note/note.model';
import noteStore from 'src/module/note/note.store';

interface INoteEditBookFormProps {
    note: INote,
    visible: boolean,
    onCancel: () => void,
    refreshNoteBoard: () => Promise<void>
}

const NoteEditBookForm: FC<INoteEditBookFormProps> = ({
    note,
    visible,
    onCancel,
    refreshNoteBoard
}) => {
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    const [background, setBackground] = useState(note.color)

    useEffect(() => {
        setBackground(note.color)
        form.setFieldsValue({
            "color": note.color
        })
    }, [note.color, form])

    const handleChangeComplete = (color: any) => {
        setBackground(color.hex)
        form.setFieldsValue({
            "color": color.hex
        })
    }

    const handleOk = () => {
        form
            .validateFields()
            .then(async (values: any) => {
                if (note.title !== values.title || note.color !== values.color) {
                    await noteStore.updateNoteBrief(note.id, values.title, values.color)
                    await refreshNoteBoard()
                }
                form.resetFields()
                onCancel()
            })
            .catch(info => {
                console.log('Validate Failed:', info);
            })
    }
    return (
        <Modal
            cancelText="取消"
            visible={visible}
            title={"修改"}
            okText="确认"
            onCancel={onCancel}
            onOk={handleOk}
            width={450}
            getContainer={false}
            destroyOnClose={false}
            mask={true}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="deviceDashboardUpdateForm"
                style={{ maxHeight: 600 }}
                initialValues={{ title: "快速笔记", color: note.color }}>
                <Form.Item
                    name="title"
                    label="标题"
                    rules={[{ required: true, message: '请输入标题!' }, {
                        max: 50, message: '标题最多50字'
                    }]}
                >
                    <Input placeholder="标题" />
                </Form.Item>
                <Form.Item
                    name="color"
                    label="颜色"
                >
                    <SketchPicker
                        width="200"
                        disableAlpha={true}
                        color={background}
                        onChangeComplete={handleChangeComplete}
                    />
                </Form.Item>
            </Form>
        </Modal >
    )
}

export default NoteEditBookForm
