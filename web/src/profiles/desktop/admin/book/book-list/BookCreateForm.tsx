import { Modal, Form, Input, InputNumber, Select } from 'antd'
import { DatePicker, Uploader } from 'src/components'
import { FC, useState } from 'react'
import { IBook } from 'src/module/book/book.model'

interface IBookCreateFormProps {
    visible: boolean
    onCreate: (values: IBook) => void
    onCancel: () => void
}

export const BookCreateForm: FC<IBookCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [url, setUrl] = useState('')
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    return (
        <Modal
            visible={visible}
            title="新增书籍"
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
                        onCreate(values)
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
                name="bookCreateForm"
                style={{ maxHeight: 600, overflowY: "auto" }}
                initialValues={{}}
            >
                <Form.Item
                    name="isbn"
                    label="isbn"
                    rules={[{ required: true, message: '请输入isbn!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="name"
                    label="书名"
                    rules={[{ required: true, message: '请输入书名!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="desc"
                    label="简介"
                    rules={[{ required: true, message: '请输入简介!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="cover"
                    label="封面"
                >
                    <Uploader
                        fileLimit={1}
                        bucketName="image"
                        validFileTypes={["image/jpeg", "image/png", "image/webp"]}
                        setURL={setUrl}
                    />
                </Form.Item>
                <Form.Item
                    name="author"
                    label="作者"
                    rules={[{ required: true, message: '请输入作者!' }]}
                >
                    <Select
                        mode="tags"
                        size={"large"}
                        style={{ width: '100%' }}
                    >
                    </Select>
                </Form.Item>
                <Form.Item
                    name="translator"
                    label="译者"
                    rules={[{ required: true, message: '请输入译者!' }]}
                >
                    <Select
                        mode="tags"
                        size={"large"}
                        style={{ width: '100%' }}
                    >
                    </Select>
                </Form.Item>
                <Form.Item
                    name="publishingHouse"
                    label="出版社"
                    rules={[{ required: true, message: '请输入出版社!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="edition"
                    label="版次"
                    rules={[{ required: true, message: '请输入版次!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="printedTimes"
                    label="印次"
                    rules={[{ required: true, message: '请输入印次!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="printedSheets"
                    label="印张"
                    rules={[{ required: true, message: '请输入印张!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="format"
                    label="开本"
                    rules={[{ required: true, message: '请输入开本!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="wordCount"
                    label="字数"
                    rules={[{ required: true, message: '请输入字数!' }]}
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item
                    name="pricing"
                    label="定价"
                    rules={[{ required: true, message: '请输入定价!' }]}
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item
                    name="purchasePrice"
                    label="购买价"
                    rules={[{ required: true, message: '请输入购买价!' }]}
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item
                    name="purchaseTime"
                    label="购买时间"
                    rules={[{ required: true, message: '请输入购买时间!' }]}
                >
                    <DatePicker />
                </Form.Item>
                <Form.Item
                    name="purchaseSource"
                    label="购买途径"
                    rules={[{ required: true, message: '请输入购买途径!' }]}
                >
                    <Input />
                </Form.Item>
            </Form>
        </Modal>
    )
}
