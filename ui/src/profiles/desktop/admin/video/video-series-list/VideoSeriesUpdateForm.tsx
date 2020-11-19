import { Modal, Form, Input } from 'antd'
import React, { FC, useEffect } from 'react'

interface IUpdateVideoProps {
    id: number,
    name: string,
}
interface IVideoUpdateSeriesFormProps {
    visible: boolean
    data: IUpdateVideoProps,
    onUpdate: (values: IUpdateVideoProps) => void
    onCancel: () => void
}

export const VideoSeriesUpdateForm: FC<IVideoUpdateSeriesFormProps> = ({
    visible,
    data,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    useEffect(() => {
        form.setFieldsValue({
            "id": data.id,
            "name": data.name,
        })
    }, [form, data])

    return (
        <Modal
            visible={visible}
            title="编辑视频系列"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                }
            }
            getContainer={false}
            onOk={() => {
                form
                    .validateFields()
                    .then((values: any) => {
                        form.resetFields()
                        onUpdate(values)
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info)
                    })
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="videoSeriesUpdateForm"
            >
                <Form.Item
                    name="id"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input />
                </Form.Item>
            </Form>
        </Modal>
    )
}
