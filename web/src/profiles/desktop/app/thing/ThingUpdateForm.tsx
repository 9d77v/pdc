import { Modal, Form, Input, InputNumber, Select, Tooltip } from 'antd'
import { DatePicker } from 'src/components'
import { useState, useEffect, FC } from 'react'
import dayjs from 'dayjs'
import { ConsumerExpenditureMap, RubbishCategoryMap, ThingStatusMap, TagStyle } from 'src/consts/consts'
import { QuestionCircleOutlined } from '@ant-design/icons'
import { Uploader } from 'src/components'

interface Values {
    title: string
    description: string
}

interface IUpdateThingProps {
    id: number
    name: string
    num: number
    brandName: string
    pics: string[]
    unitPrice: number
    unit: string
    specifications: string
    category: number
    consumerExpenditure: string
    location: string
    status: number
    purchaseDate: number
    purchasePlatform: string
    refOrderID: string
    rubbishCategory: number[]
}
interface IThingUpdateFormProps {
    visible: boolean
    data: IUpdateThingProps,
    onUpdate: (values: Values) => void
    onCancel: () => void
}

export const ThingUpdateForm: FC<IThingUpdateFormProps> = ({
    visible,
    data,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm()
    const emptURLS: string[] = []
    const [thingURLs, setThingURLs] = useState(emptURLS)
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    useEffect(() => {
        form.setFieldsValue({
            "id": data.id,
            "name": data.name,
            "num": data.num,
            "brandName": data.brandName,
            "pics": data.pics,
            "unitPrice": data.unitPrice,
            "unit": data.unit,
            "specifications": data.specifications,
            "category": data.category,
            "consumerExpenditure": data.consumerExpenditure,
            "location": data.location,
            "status": data.status,
            "purchaseDate": dayjs(data.purchaseDate * 1000),
            "purchasePlatform": data.purchasePlatform,
            "refOrderID": data.refOrderID,
            "rubbishCategory": data.rubbishCategory
        })
        setThingURLs(data.pics)
    }, [form, data])

    let categoryOptions: any[] = []
    ConsumerExpenditureMap.forEach((value: string, key: string) => {
        categoryOptions.push(<Select.Option value={key} key={'category_options_' + key}>{value}</Select.Option>)
    })
    let rubbishCategoryOptions: any[] = []
    RubbishCategoryMap.forEach((value: TagStyle, key: number) => {
        rubbishCategoryOptions.push(<Select.Option value={key} key={'rubbish_cateogry_options_' + key}>{value.text}</Select.Option>)
    })
    let statusOptions: any[] = []
    ThingStatusMap.forEach((value: TagStyle, key: number) => {
        statusOptions.push(<Select.Option value={key} key={'thing_stauts_options_' + key}>{value.text}</Select.Option>)
    })
    return (
        <Modal
            visible={visible}
            title="编辑物品"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setThingURLs([])
                }
            }
            getContainer={false}
            onOk={() => {
                form.setFieldsValue({
                    "pics": thingURLs,
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
                setThingURLs([])
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="thingUpdateForm"
                style={{ maxHeight: 600, overflowY: 'auto' }}
                initialValues={{ num: 1, unitPrice: 0, category: "01", rubbishCategory: [0], status: 1 }}
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
                <Form.Item
                    name="num"
                    label="数量"
                    rules={[{ required: true, message: '请输入数量!' }]}
                >
                    <InputNumber min={1} step={1} />
                </Form.Item>
                <Form.Item
                    name="unitPrice"
                    label="单价"
                    rules={[{ required: true, message: '请输入单价!' }]}
                >
                    <InputNumber min={0} step={0.01} />
                </Form.Item>
                <Form.Item
                    name="consumerExpenditure"
                    label={<span>
                        消费支出&nbsp
                        <Tooltip title={<a href="http://www.stats.gov.cn/tjsj/tjbz/201310/P020131021349384303616.pdf" target="_blank" rel="noopener noreferrer">居民消费支出分类（2013）</a>}>
                            <QuestionCircleOutlined />
                        </Tooltip>
                    </span>}
                    hasFeedback
                    rules={[{ required: true, message: '请选择一个消费支出!' }]}
                >
                    <Select placeholder="请选择一个消费支出!">
                        {categoryOptions}
                    </Select>
                </Form.Item>
                <Form.Item
                    name="category"
                    label="分类"
                    hasFeedback
                    rules={[{ required: true, message: '请选择一个分类!' }]}
                    noStyle
                >
                    <Select placeholder="请选择一个分类!" style={{ display: "none" }}>
                        {categoryOptions}
                    </Select>
                </Form.Item>
                <Form.Item
                    name="purchaseDate"
                    label="购买时间"
                    rules={[{ required: true, message: '请选择购买时间!' }]}
                >
                    <DatePicker />
                </Form.Item>
                <Form.Item
                    name="rubbishCategory"
                    label="垃圾分类"
                    hasFeedback
                    rules={[{ message: '请选择垃圾分类!', type: 'array' }]}
                >
                    <Select mode="multiple" placeholder="请选择垃圾分类!">
                        {rubbishCategoryOptions}
                    </Select>
                </Form.Item>
                <Form.Item
                    name="status"
                    label="状态"
                    hasFeedback
                    rules={[{ required: true, message: '请选择状态!' }]}
                >
                    <Select placeholder="请选择状态!">
                        {statusOptions}
                    </Select>
                </Form.Item>
                <Form.Item name="pics" label="图片">
                    <Uploader
                        fileLimit={10}
                        bucketName="image"
                        validFileTypes={["image/jpeg", "image/png", "image/webp"]}
                        setURL={setThingURLs}
                    />
                </Form.Item>
                <Form.Item
                    name="brandName"
                    label="品牌名称"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="unit"
                    label="单位"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="specifications"
                    label="规格"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="location"
                    label="位置"
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    name="purchasePlatform"
                    label="购买平台"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="refOrderID"
                    label="关联订单号"
                >
                    <Input />
                </Form.Item>
            </Form>
        </Modal>
    )
}
