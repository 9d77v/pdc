import { Modal, Form, Input, DatePicker, InputNumber, Select } from 'antd';
import React, { useState } from 'react'
import { Uploader } from '../../../components/Uploader';
import { CategoryMap, RubbishCategoryMap, TagStyle, ThingStatusMap } from '../../../consts/category_data';
import moment from 'moment';

interface Values {
    title: string;
    description: string;
}

interface ThingCreateFormProps {
    visible: boolean;
    onCreate: (values: Values) => void;
    onCancel: () => void;
}

export const ThingCreateForm: React.FC<ThingCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [thingURLs, setThingURLs] = useState([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    let categoryOptions: any[] = []
    CategoryMap.forEach((value: string, key: string) => {
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
            title="新增物品"
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
                        form.resetFields();
                        onCreate(values);
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info);
                    });
                setThingURLs([])
            }}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="thingCreateForm"
                style={{ maxHeight: 600, overflowY: 'scroll' }}
                initialValues={{ num: 1, unitPrice: 0, category: "01", rubbishCategory: [0], status: 1, purchaseDate: moment(0) }}
            >
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
                    name="category"
                    label="类别"
                    hasFeedback
                    rules={[{ required: true, message: '请选择一个类别!' }]}

                >
                    <Select placeholder="请选择一个类别!">
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
    );
};