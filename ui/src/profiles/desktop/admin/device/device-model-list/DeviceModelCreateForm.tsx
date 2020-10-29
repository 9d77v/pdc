import { Modal, Form, Input, Select } from 'antd';
import React from 'react'
import TextArea from 'antd/lib/input/TextArea';
import { DeviceTypeMap } from 'src/consts/consts';


export interface INewDeviceModel {
    name: string
    deviceType: number
    desc: string
}

interface DeviceModelCreateFormProps {
    visible: boolean;
    onCreate: (values: INewDeviceModel) => void;
    onCancel: () => void;
}

export const DeviceModelCreateForm: React.FC<DeviceModelCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    let deviceTypeOptions: any[] = []
    DeviceTypeMap.forEach((value: string, key: number) => {
        deviceTypeOptions.push(<Select.Option
            value={key}
            key={'deviceType_options_' + key}>{value}</Select.Option>)
    })
    return (
        <Modal
            visible={visible}
            title="新增设备模型"
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
                name="deviceModelCreateForm"
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
                {/* <Form.Item
                    name="deviceType"
                    label="设备类型"
                    noStyle
                    hasFeedback
                    rules={[{ required: true, message: '请选择设备类型!' }]}
                >
                    <Select >
                        {deviceTypeOptions}
                    </Select>
                </Form.Item> */}
                <Form.Item name="desc" label="简介">
                    <TextArea rows={4} />
                </Form.Item>
            </Form>
        </Modal>
    );
};