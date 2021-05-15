import { Modal, Form, Input, Select } from 'antd'
import { useState, useEffect, FC } from 'react'
import { useQuery } from '@apollo/react-hooks'
import { Img } from 'src/components'
import { BOOK_COMBO } from 'src/gqls/book/book.query'
import { IBookPosition } from 'src/module/book/book_position.model'

interface IAddBookFormProps {
    visible: boolean
    onCreate: (values: any) => void
    onCancel: () => void
    addBookData: IBookPosition
}

const { Option } = Select
export const AddBookForm: FC<IAddBookFormProps> = ({
    visible,
    onCreate,
    onCancel,
    addBookData
}) => {
    const [form] = Form.useForm()
    const [value, setValue] = useState<number[]>([])
    const [keyword, setKeyword] = useState("")

    const { data, refetch } = useQuery(BOOK_COMBO,
        {
            variables: {
                searchParam: {
                    page: 1,
                    pageSize: 30,
                    keyword: keyword,
                    sorts: [{
                        field: 'id',
                        isAsc: false
                    }]
                }
            },
            fetchPolicy: "cache-and-network"
        })

    const onFinish = (values: any) => {
        console.log('Finish:', values)
    }

    useEffect(() => {
        if (visible) {
            refetch()
        }
    }, [visible, refetch])

    useEffect(() => {
        form.setFieldsValue({
            "bookshelfID": addBookData.bookshelf_id,
            "layer": addBookData.layer,
            "partition": addBookData.partition
        })
    }, [form, addBookData])

    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }

    let timer: any
    const handleSearch = (value: string) => {
        clearTimeout(timer)
        timer = setTimeout(() => {
            setKeyword(value)
        }, 1000)
    }

    const handleChange = (v: number) => {
        setValue([...value, v])
    }
    const options = data === undefined ? null : data.books.edges.map((d: any) =>
        <Option key={d.value} value={d.value}>
            <div style={{ display: "flex", height: 30 }}><Img src={d?.cover} width={60} height={30} />
                <div style={{ display: "flex", alignItems: "center", justifyContent: "center", padding: 10 }}>{d.text}</div></div></Option>)
    return (
        <Modal
            visible={visible}
            title="放入书籍"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields(["bookIDs", "layer", "partition"])
                }
            }
            getContainer={false}
            onOk={() => {
                form
                    .validateFields()
                    .then((values: any) => {
                        onCreate(values)
                        form.resetFields(["bookIDs", "layer", "partition"])
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
                name="addBookCreateForm"
                onFinish={onFinish}
            >
                <Form.Item
                    name="bookshelfID"
                    label="书架"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="bookIDs"
                    label="书本"
                    rules={[{ required: true, message: '请选择书本!' }]}
                >
                    <Select
                        showSearch
                        defaultActiveFirstOption={false}
                        mode="multiple"
                        allowClear
                        showArrow={true}
                        filterOption={false}
                        onSearch={handleSearch}
                        onChange={handleChange}
                        notFoundContent={null}
                    >
                        {options}
                    </Select>
                </Form.Item>
                <Form.Item
                    name="layer"
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="partition"
                >
                    <Input hidden />
                </Form.Item>
            </Form>
        </Modal >
    )
}
