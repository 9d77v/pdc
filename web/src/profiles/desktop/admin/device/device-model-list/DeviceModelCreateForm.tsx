import { Modal, Form, Input, Select } from 'antd';
import { FC, useState } from 'react'
import TextArea from 'antd/lib/input/TextArea';
import { CameraCompanyMap, DeviceTypeMap } from 'src/consts/consts';
import { INewDeviceModel } from 'src/models/device';

interface IDeviceModelCreateFormProps {
    visible: boolean;
    onCreate: (values: INewDeviceModel) => void;
    onCancel: () => void;
}

export const DeviceModelCreateForm: FC<IDeviceModelCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [hideCameraCompanyOption, setHideCameraCompanyOption] = useState(true)
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    let deviceTypeOptions: any[] = []
    DeviceTypeMap.forEach((value: string, key: number) => {
        deviceTypeOptions.push(<Select.Option
            value={key}
            key={'deviceType_options_' + key}>{value}</Select.Option>)
        return
    })
    let cameraCompanyOptions: any = []
    CameraCompanyMap.forEach((value: string, key: number) => {
        cameraCompanyOptions.push(<Select.Option
            value={key}
            key={'cameraCompany_options_' + key}>{value}</Select.Option>)
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
                initialValues={{ deviceType: 0, cameraCompany: 0 }}
            >
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item name="desc" label="简介">
                    <TextArea rows={4} />
                </Form.Item>
                <Form.Item
                    name="deviceType"
                    label="设备类型"
                    hidden
                    hasFeedback
                    rules={[{ required: true, message: '请选择设备类型!' }]}
                >
                    <Select onChange={(value) => {
                        setHideCameraCompanyOption(Number(value) !== 1)
                    }} >
                        {deviceTypeOptions}
                    </Select>
                </Form.Item>
                <Form.Item
                    name="cameraCompany"
                    label="摄像头厂家"
                    hidden={hideCameraCompanyOption}
                    hasFeedback
                    rules={[{ required: true, message: '请选择设备类型!' }]}
                >
                    <Select >
                        {cameraCompanyOptions}
                    </Select>
                </Form.Item>
            </Form>
        </Modal>
    );
};
