import { Form, Radio } from 'antd'
import React, { forwardRef, Ref, useImperativeHandle, useState } from 'react'
import { Uploader } from 'src/components/Uploader'
import { supportedSubtitleTypes } from 'src/consts/consts'

const VideoCreateStepThreeForm = (props: any, ref: Ref<any>) => {
    const [form] = Form.useForm()
    const getForm = () => {
        return form
    }
    const getSubtitles = () => {
        return subtitles
    }
    const resetSubtitles = () => {
        setSubtitles([])
    }
    useImperativeHandle(ref, () => ({
        getForm,
        getSubtitles,
        resetSubtitles
    }))

    const [subtitles, setSubtitles] = useState([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }

    return (
        <Form
            {...layout}
            form={form}
            layout="horizontal"
            name="videoCreateForm"
            initialValues={{ subtitle_lang: "简体中文" }}
        >
            <Form.Item name="subtitle_lang" label="字幕语言">
                <Radio.Group buttonStyle="solid">
                    <Radio.Button value="简体中文">简体中文</Radio.Button>
                    <Radio.Button value="中日双语">中日双语</Radio.Button>
                </Radio.Group>
            </Form.Item>
            <Form.Item name="subtitles" label="字幕列表">
                <Uploader
                    fileLimit={0}
                    bucketName="vtt"
                    validFileTypes={supportedSubtitleTypes}
                    setURL={setSubtitles}
                />
            </Form.Item>
        </Form>
    )
}

export default forwardRef(VideoCreateStepThreeForm)
