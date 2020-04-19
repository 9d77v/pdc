import { Modal, Form, Input, InputNumber, Button, Typography, Radio } from 'antd';
import React, { useState, useRef, useEffect } from 'react'
import { Uploader } from '../../components/Uploader';
interface IEpisode {
    title?: string;
    description?: string;
    num: number,
}

interface EpisodeUpdateFormProps {
    visible: boolean;
    onUpdate: (values: IEpisode) => void;
    onCancel: () => void;
    values: IEpisode,
}

interface ModalFormProps {
    visible: boolean;
    onCancel: () => void;
}

// reset form fields when modal is form, closed
const useResetFormOnCloseModal = ({ form, visible }: any) => {
    const prevVisibleRef = useRef();
    useEffect(() => {
        prevVisibleRef.current = visible;
    }, [visible]);
    const prevVisible = prevVisibleRef.current;

    useEffect(() => {
        if (!visible && prevVisible) {
            form.resetFields();
        }
    }, [form, prevVisible, visible]);
};

const ModalSubtitleForm: React.FC<ModalFormProps> = ({ visible, onCancel }) => {
    const [form] = Form.useForm();

    const [url, setUrl] = useState("")
    useResetFormOnCloseModal({
        form,
        visible,
    });

    const onOk = () => {
        form.setFieldsValue({
            "url": url
        })
        form.submit();
    };

    return (
        <Modal title="新增字幕" visible={visible}
            onOk={onOk}
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setUrl('')
                }
            }>
            <Form form={form} layout="vertical" name="subtitleForm" initialValues={{ name: "简体中文" }}>
                <Form.Item name="name" label="标签" rules={[{ required: true }]}>
                    <Radio.Group buttonStyle="solid">
                        <Radio.Button value="简体中文">简体中文</Radio.Button>
                        <Radio.Button value="中日双语">中日双语</Radio.Button>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="url" label="地址" rules={[{ required: true }]}>
                    <Uploader
                        fileLimit={1}
                        bucketName="vtt"
                        validFileTypes={["text/vtt"]}
                        setURL={setUrl}
                    />
                </Form.Item>
            </Form>
        </Modal>
    );
};

export const EpisodeUpdateForm: React.FC<EpisodeUpdateFormProps> = ({
    visible,
    onUpdate,
    onCancel,
    values
}) => {
    const [form] = Form.useForm();
    const [url, setUrl] = useState('')

    const [subtitleVisible, setSubtitleVisible] = useState(false);

    const showSubtitleModal = () => {
        setSubtitleVisible(true);
    };

    const hideSubtitleModal = () => {
        setSubtitleVisible(false);
    };

    const onFinish = (values: any) => {
        console.log('Finish:', values);
    };

    return (
        <Modal
            visible={visible}
            title="新增分集"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setUrl('')
                }
            }
            onOk={() => {
                form.setFieldsValue({
                    "url": url
                })
                form
                    .validateFields()
                    .then((values: any) => {
                        const subtitles = form.getFieldValue('subtitles') || [];
                        values.subtitles = subtitles
                        onUpdate(values);
                        form.resetFields();
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info);
                    });
                setUrl('')
            }}
        >
            <Form.Provider
                onFormFinish={(name, { values, forms }) => {
                    if (name === 'subtitleForm') {
                        const { basicForm } = forms;
                        const subtitles = basicForm.getFieldValue('subtitles') || [];
                        basicForm.setFieldsValue({ subtitles: [...subtitles, values] });
                        setSubtitleVisible(false);
                    }
                }}
            >
                <Form
                    form={form}
                    layout="vertical"
                    name="basicForm"
                    onFinish={onFinish}
                    initialValues={{ num: values.num, title: values.title }}
                >
                    <Form.Item
                        name="title"
                        label="标题"
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item name="num" label="话"
                        rules={[{ required: true, message: '请输入话数!' }]}
                    >
                        <InputNumber />
                    </Form.Item>
                    <Form.Item name="desc" label="简介">
                        <Input type="textarea" />
                    </Form.Item>
                    <Form.Item
                        name="url"
                        label="上传视频"
                        rules={[{ required: true, message: '请上传视频!' }]}
                    >
                        <Uploader
                            fileLimit={1}
                            bucketName="video"
                            validFileTypes={["video/mp4"]}
                            setURL={setUrl}
                        />
                    </Form.Item>
                    <Form.Item
                        label="字幕列表"
                        shouldUpdate={(prevValues, curValues) => prevValues.subtitles !== curValues.subtitles}
                    >
                        {({ getFieldValue }) => {
                            const subtitles = getFieldValue('subtitles') || [];
                            return subtitles.length ? (
                                <ul>
                                    {subtitles.map((subtitle: any, index: number) => (
                                        <li key={index}>
                                            {subtitle.name} - {subtitle.url}
                                        </li>
                                    ))}
                                </ul>
                            ) : (
                                    <Typography.Text className="ant-form-text" type="secondary">
                                        暂无字幕
                                    </Typography.Text>
                                );
                        }}
                    </Form.Item>

                    <Button htmlType="button" style={{ margin: '0 8px' }} onClick={showSubtitleModal}>
                        添加字幕
                </Button>
                </Form>
                <ModalSubtitleForm visible={subtitleVisible} onCancel={hideSubtitleModal} />
            </Form.Provider>
        </Modal >
    );
};