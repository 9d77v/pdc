import { Modal, Form, Input } from 'antd'
import React, { useEffect } from 'react'

interface VideoSeriesItemUpdateFormProps {
    visible: boolean
    onUpdate: (values: any) => void
    onCancel: () => void
    data: any
}

export const VideoSeriesItemUpdateForm: React.FC<VideoSeriesItemUpdateFormProps> = ({
    visible,
    onUpdate,
    onCancel,
    data
}) => {
    const [form] = Form.useForm()

    const onFinish = (values: any) => {
        console.log('Finish:', values)
    }

    useEffect(() => {
        form.setFieldsValue({
            "video_id": data.video_id,
            "video_series_id": data.video_series_id,
            "alias": data.alias,
        })
    }, [data, form])

    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    return (
        <Modal
            visible={visible}
            title="修改视频"
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
                        onUpdate(values)
                        form.resetFields()
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
                name="videoSeriesItemUpdateForm"
                onFinish={onFinish}
            >
                <Form.Item
                    name="video_series_id"
                    label="视频系列"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="video_id"
                    label="视频"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="alias"
                    label="别名"
                    rules={[{ required: true, message: '请设置视频别名!' }]}
                >
                    <Input />
                </Form.Item>
            </Form>
        </Modal >
    )
}