import React, { useState, useRef, useEffect } from 'react'
import { Modal, Form, Radio } from 'antd'
import { Uploader } from '../../../../components/Uploader'

interface ModalFormProps {
    visible: boolean
    onCancel: () => void
}

// reset form fields when modal is form, closed
const useResetFormOnCloseModal = ({ form, visible }: any) => {
    const prevVisibleRef = useRef()
    useEffect(() => {
        prevVisibleRef.current = visible
    }, [visible])
    const prevVisible = prevVisibleRef.current

    useEffect(() => {
        if (!visible && prevVisible) {
            form.resetFields()
        }
    }, [form, prevVisible, visible])
}

export const SubtitleForm: React.FC<ModalFormProps> = ({ visible, onCancel }) => {
    const [form] = Form.useForm()

    const [url, setUrl] = useState("")
    useResetFormOnCloseModal({
        form,
        visible,
    })

    const onOk = () => {
        form.setFieldsValue({
            "url": url
        })
        form.submit()
    }

    return (
        <Modal title="新增字幕" visible={visible} onOk={onOk}
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setUrl('')
                }
            }        >
            <Form form={form} layout="vertical" name="subtitleForm" initialValues={{ name: "简体中文" }}>
                <Form.Item name="name" label="标签" rules={[{ required: true }]}>
                    <Radio.Group buttonStyle="solid">
                        <Radio.Button value="简体中文">简体中文</Radio.Button>
                        <Radio.Button value="中日双语">中日双语</Radio.Button>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="url" label="字幕（支持上传ass或srt转vtt格式）" rules={[{ required: true }]}>
                    <Uploader
                        fileLimit={1}
                        bucketName="vtt"
                        validFileTypes={["text/vtt", "text/ass", 'text/srt']}
                        setURL={setUrl}
                    />
                </Form.Item>
            </Form>
        </Modal>
    )
}