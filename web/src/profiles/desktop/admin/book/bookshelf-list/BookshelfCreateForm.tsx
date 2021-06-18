import { Modal, Form, Input, InputNumber } from 'antd'
import { FC, useState } from 'react'
import { IBookshelf } from 'src/module/book/bookshelf.model'
import { Uploader } from 'src/components'

interface IBookshelfCreateFormProps {
    visible: boolean
    onCreate: (values: IBookshelf) => void
    onCancel: () => void
}

export const BookshelfCreateForm: FC<IBookshelfCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [url, setUrl] = useState<string[]>([])
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    return (
        <Modal
            visible={visible}
            title="新增书架"
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
                    "cover": url[0]
                })
                form
                    .validateFields()
                    .then((values: any) => {
                        form.resetFields()
                        onCreate(values)
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info)
                    })
                setUrl([])
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="bookshelfCreateForm"
                style={{ maxHeight: 600, overflowY: "auto" }}
                initialValues={{}}
            >
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
                    rules={[{ required: true, message: '请上传图片!' }]}
                >
                    <Uploader
                        fileLimit={1}
                        bucketName="image"
                        validFileTypes={["image/jpeg", "image/png", "image/webp"]}
                        setURL={setUrl}
                    />
                </Form.Item>
                <Form.Item
                    name="layerNum"
                    label="层数"
                    rules={[{ required: true, message: '请输入层数!' }]}
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item
                    name="partitionNum"
                    label="分区数"
                    rules={[{ required: true, message: '请输入分区数!' }]}
                >
                    <InputNumber />
                </Form.Item>
            </Form>
        </Modal>
    )
}
