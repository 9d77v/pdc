import { Modal, Form, Input } from 'antd';
import { FC, useEffect, useState } from 'react';
import { SketchPicker } from 'react-color'

interface INoteBookNewFormProps {
    visible: boolean
    newClolor: string
    addNewNotebook: (values: any) => Promise<void>
    onCancel: () => void
}

export const NoteBookNewForm: FC<INoteBookNewFormProps> = ({
    visible,
    newClolor,
    addNewNotebook,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    const [background, setBackground] = useState(newClolor)

    useEffect(() => {
        setBackground(newClolor)
        form.setFieldsValue({
            "color": newClolor
        })
    }, [newClolor, form])

    const handleOk = () => {
        form
            .validateFields()
            .then(async (values: any) => {
                form.resetFields()
                onCancel()
                await addNewNotebook(values)
            })
            .catch(info => {
                console.log('Validate Failed:', info);
            });
    }

    const handleCancel = () => {
        form.resetFields()
        onCancel()
    }

    const handleChangeComplete = (color: any) => {
        setBackground(color.hex)
        form.setFieldsValue({
            "color": color.hex
        })
    }
    return (
        <Modal
            visible={visible}
            title={"新建文件夹"}
            okText="确认"
            onCancel={handleCancel}
            onOk={handleOk}
            width={450}
            destroyOnClose={false}
            mask={true}
            cancelText="取消"
            getContainer={false}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="deviceDashboardCreateForm"
                style={{ maxHeight: 600 }}
                initialValues={{ title: "快速笔记", color: newClolor }}>
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
        </Modal>
    )
}
