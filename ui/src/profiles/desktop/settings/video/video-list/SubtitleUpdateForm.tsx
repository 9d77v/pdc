import { Modal, Form, Input, Radio } from 'antd';
import React, { useState } from 'react'
import { Uploader } from '../../../../../components/Uploader';

interface SubtitleUpdateFormProps {
    visible: boolean;
    onUpdate: (values: any) => void;
    onCancel: () => void;
}

export const SubtitleUpdateForm: React.FC<SubtitleUpdateFormProps> = ({
    visible,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [subtitles, setSubtitles] = useState([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }

    const onOK = () => {
        form.setFieldsValue({
            "subtitles": subtitles
        })
        form
            .validateFields()
            .then((values: any) => {
                form.resetFields();
                onUpdate(values);
            })
            .catch(info => {
                console.log('Validate Failed:', info);
            })
    }
    return (
        <Modal
            visible={visible}
            title="更换字幕"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setSubtitles([])
                }
            }
            getContainer={false}
            onOk={() => {
                if (subtitles.length === 0) {
                    Modal.confirm({
                        title: "当前字幕为空，确认要清空字幕吗",
                        onOk: onOK
                    })
                } else {
                    onOK()
                }
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="subtitleUpdateForm"
                initialValues={{ subtitle_lang: "简体中文" }}
            >
                <Form.Item
                    name="id"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item name="subtitle_lang" label="字幕语言"
                    rules={[{ required: true, message: '请选择字幕语言!' }]}
                >
                    <Radio.Group buttonStyle="solid">
                        <Radio.Button value="简体中文">简体中文</Radio.Button>
                        <Radio.Button value="中日双语">中日双语</Radio.Button>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="subtitles" label="字幕列表">
                    <Uploader
                        fileLimit={0}
                        bucketName="vtt"
                        validFileTypes={["text/vtt", "text/ass", 'text/srt']}
                        setURL={setSubtitles}
                    />
                </Form.Item>
            </Form>
        </Modal>
    )
}
