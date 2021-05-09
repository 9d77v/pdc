import { Modal, Form, Input, InputNumber } from 'antd'
import { FC, useEffect, useState } from 'react'
import { Uploader } from 'src/components'
import { IUpdateBookshelf } from 'src/module/book/bookshelf.model'

interface IBookshelfUpdateFormProps {
    visible: boolean
    data: IUpdateBookshelf,
    onUpdate: (values: IUpdateBookshelf) => void
    onCancel: () => void
}

export const BookshelfUpdateForm: FC<IBookshelfUpdateFormProps> = ({
    visible,
    data,
    onUpdate,
    onCancel,
}) => {
    const [url, setUrl] = useState('')
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    useEffect(() => {
        form.setFieldsValue({
            "id": data.id,
            "name": data.name,
            "cover": data.cover,
            "layerNum": data.layerNum,
            "partitionNum": data.partitionNum,
        })
    }, [form, data])
    return (
        <Modal
            visible={visible}
            title="编辑书架"
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
                form.setFieldsValue({
                    "cover": url
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
                setUrl('')
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="bookshelfUpdateForm"
                style={{ maxHeight: 600, overflowY: "auto" }}
            >
                <Form.Item
                    name="id"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="name"
                    label="书架名"
                    rules={[{ required: true, message: '请输入书架名!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="cover"
                    label="图片"
                >
                    <Uploader
                        fileLimit={1}
                        bucketName="image"
                        validFileTypes={["image/jpeg", "image/png", "image/webp"]}
                        setURL={setUrl}
                    />
                </Form.Item>
            </Form>
        </Modal>
    )
}
