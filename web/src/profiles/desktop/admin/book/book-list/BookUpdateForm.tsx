import { Modal, Form, Input, InputNumber, Select, message } from 'antd'
import { DatePicker } from 'src/components'
import { FC, useEffect, useState } from 'react'
import { IUpdateBook } from 'src/module/book/book.model'
import { BOOK_DETAIL } from 'src/gqls/book/book.query'
import { useQuery } from '@apollo/react-hooks'
import dayjs from 'dayjs'
import { Uploader } from 'src/components'

interface IBookUpdateFormProps {
    visible: boolean
    id: number,
    onUpdate: (values: IUpdateBook) => void
    onCancel: () => void
}
const { TextArea } = Input;

export const BookUpdateForm: FC<IBookUpdateFormProps> = ({
    visible,
    id,
    onUpdate,
    onCancel,
}) => {
    const [url, setUrl] = useState<string[]>([])
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    const { error, data } = useQuery(BOOK_DETAIL,
        {
            variables: {
                searchParam: {
                    ids: [id]
                },
            },
            fetchPolicy: "cache-and-network"
        })

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    useEffect(() => {
        const book = data?.books?.edges[0]
        if (book) {
            form.setFieldsValue({
                "id": book.id,
                "isbn": book.isbn,
                "name": book.name,
                "desc": book.desc,
                "cover": book.cover,
                "author": book.author,
                "translator": book.translator,
                "publishingHouse": book.publishingHouse,
                "edition": book.edition,
                "printedTimes": book.printedTimes,
                "printedSheets": book.printedSheets,
                "format": book.format,
                "wordCount": book.wordCount,
                "pricing": book.pricing,
                "purchasePrice": book.purchasePrice,
                "purchaseTime": book.purchaseTime ? dayjs(book.purchaseTime * 1000) : undefined,
                "purchaseSource": book.purchaseSource,
                "bookBorrowUID": book.bookBorrowUID,
            })
        }
    }, [form, data])

    return (
        <Modal
            visible={visible}
            title="编辑书籍"
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
                        onUpdate(values)

                        form.resetFields()
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
                name="bookUpdateForm"
                style={{ maxHeight: 600, overflowY: "auto" }}
            >
                <Form.Item
                    name="id"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
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
                    <TextArea />
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
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="printedTimes"
                    label="印次"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="printedSheets"
                    label="印张"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="format"
                    label="开本"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="wordCount"
                    label="字数"
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
                    name="packing"
                    label="包装"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="pageSize"
                    label="页数"
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item
                    name="purchasePrice"
                    label="购买价"
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
