import { Modal, Form } from 'antd'
import { useState, FC } from 'react'
import { Uploader } from 'src/components'

interface Values {
    title: string
    description: string
}

interface IVideoResourceUpdateFormProps {
    visible: boolean
    onUpdate: (values: Values) => void
    onCancel: () => void
}

export const VideoResourceUpdateForm: FC<IVideoResourceUpdateFormProps> = ({
    visible,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm()
    const [url, setUrl] = useState<string[]>([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }

    const handleOK = () => {
        form.setFieldsValue({
            "urls": url,
        })
        form
            .validateFields()
            .then((values: any) => {
                form.resetFields()
                onUpdate(values)
            })
            .catch(info => {
                console.log('Validate Failed:', info)
            })
        setUrl([])
    }

    const handleCancel = () => {
        onCancel()
        form.resetFields()
        setUrl([])
    }

    return (
        <Modal
            visible={visible}
            title="更新视频源"
            okText="确定"
            cancelText="取消"
            onCancel={handleCancel}
            getContainer={false}
            onOk={handleOK}
            maskClosable={false}
            mask={true}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="videoResourceUpdateForm"
            >
                <Form.Item
                    name="videoURLs"
                    label="视频列表"
                    rules={[{ required: true, message: '请上传视频!' }]}
                >
                    <Uploader
                        fileLimit={0}
                        bucketName="video"
                        validFileTypes={["video/mp4"]}
                        setURL={setUrl}
                    />
                </Form.Item>
            </Form>
        </Modal>
    )
}
