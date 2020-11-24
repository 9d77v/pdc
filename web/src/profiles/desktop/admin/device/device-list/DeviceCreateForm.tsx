import { Modal, Form, Input, Select, InputNumber } from 'antd';
import React, { FC, useState } from 'react'
import { useQuery } from '@apollo/react-hooks';
import { DEVICE_MODEL_COMBO } from 'src/consts/device.gql';
import { INewDevice } from 'src/models/device';

interface IDeviceCreateFormProps {
    visible: boolean;
    onCreate: (values: INewDevice) => void;
    onCancel: () => void;
}
const { Option } = Select;

export const DeviceCreateForm: FC<IDeviceCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {

    const [form] = Form.useForm();
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    const [keyword, setKeyword] = useState("")
    const [value, setValue] = useState(0)
    const { data } = useQuery(DEVICE_MODEL_COMBO,
        {
            variables: {
                page: 1,
                pageSize: 10,
                keyword: keyword,
                sorts: [{
                    field: 'id',
                    isAsc: false
                }]
            },
            fetchPolicy: "cache-and-network"
        })
    let timer: any
    const handleSearch = (value: string) => {
        clearTimeout(timer)
        timer = setTimeout(() => {
            setKeyword(value)
        }, 1000)
    }

    const handleChange = (value: number) => {
        setValue(value)
    }
    const options = data === undefined ? null : data.deviceModels.edges.map((d: any) =>
        <Option key={d.value} value={d.value}>{d.text}</Option>);
    return (
        <Modal
            visible={visible}
            title="新增设备"
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
                        form.resetFields();
                        onCreate(values);
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info);
                    });
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="deviceCreateForm"
                style={{ maxHeight: 600 }}
                initialValues={{ deviceType: 0 }}
            >
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="ip"
                    label="IP地址"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="port"
                    label="端口"
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item
                    name="deviceModelID"
                    label="设备模型"
                    rules={[{ required: true, message: '请选择设备模型!' }]}
                >
                    <Select
                        showSearch
                        value={value}
                        defaultActiveFirstOption={false}
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
                    name="username"
                    label="用户名"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="password"
                    label="密码"
                >
                    <Input.Password />
                </Form.Item>
            </Form>
        </Modal>
    )
}
