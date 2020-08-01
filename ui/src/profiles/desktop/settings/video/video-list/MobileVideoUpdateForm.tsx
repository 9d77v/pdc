import { Modal, Form } from 'antd';
import React, { useState } from 'react'
import { Uploader } from '../../../../../components/Uploader';

interface MobileVideoUpdateProps {
    visible: boolean;
    videoID: number
    onUpdate: (values: any) => void;
    onCancel: () => void;
}

export const MobileVideoUpdateForm: React.FC<MobileVideoUpdateProps> = ({
    visible,
    videoID,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [videoURLs, setVideoURLs] = useState([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    const videoPathPrefix = "mobile/" + videoID.toString() + "/"
    return (
        <Modal
            visible={visible}
            title="更换视频"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setVideoURLs([])
                }
            }
            getContainer={false}
            onOk={() => {
                form.setFieldsValue({
                    "videoURLs": videoURLs,
                })
                form
                    .validateFields()
                    .then((values: any) => {
                        form.resetFields();
                        onUpdate(values);
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info);
                    });
                setVideoURLs([])
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="mobileVideoUpdateForm"
            >
                <Form.Item
                    name="videoURLs"
                    label="手机视频"
                    rules={[{ required: true, message: '请上传视频!' }]}
                >
                    <Uploader
                        fileLimit={0}
                        bucketName="video"
                        filePathPrefix={videoPathPrefix}
                        validFileTypes={["video/mp4"]}
                        setURL={setVideoURLs}
                    />
                </Form.Item>
            </Form>
        </Modal>
    );
};