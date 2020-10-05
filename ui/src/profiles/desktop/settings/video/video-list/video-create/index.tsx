import { Button, message, Steps } from 'antd';
import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import "./index.less"
import { VideoCreateStepOneForm } from './VideoCreateStepOneForm';
import { VideoCreateStepThreeForm } from './VideoCreateStepThreeForm';
import { VideoCreateStepTwoForm } from './VideoCreateStepTwoForm';
const { Step } = Steps;




export default function VideoCreateIndex() {
    const [current, setCurrent] = useState(0)
    const history = useHistory()
    const next = () => {
        setCurrent(current + 1)
    }

    const prev = () => {
        setCurrent(current - 1)
    }

    const steps = [
        {
            title: '基本信息',
            content: <VideoCreateStepOneForm />,
        },
        {
            title: '上传视频',
            content: <VideoCreateStepTwoForm />,
        },
        {
            title: '上传字幕',
            content: <VideoCreateStepThreeForm />,
        },
    ];
    return (
        <div style={{ display: "flex", flexDirection: "column", padding: 16 }}>
            <Button type="primary"
                style={{ width: 100, margin: 6 }}
                onClick={() => history.goBack()}>
                返回
                    </Button>
            < Steps current={current} >
                {
                    steps.map(item => (
                        <Step key={item.title} title={item.title} />
                    ))
                }
            </Steps >
            <div className="steps-content">{steps[current].content}</div>
            <div className="steps-action">
                {current < steps.length - 1 && (
                    <Button type="primary" style={{ float: "right" }}
                        onClick={() => next()}>
                        下一步
                    </Button>
                )}
                {current === steps.length - 1 && (
                    <Button type="primary" style={{ float: "right" }}
                        onClick={() => message.success('Processing complete!')}>
                        完成
                    </Button>
                )}
                {current > 0 && (
                    <Button style={{ margin: '0 8px', float: "right" }} onClick={() => prev()}>
                        上一步
                    </Button>
                )}
            </div>
        </div >
    )
}