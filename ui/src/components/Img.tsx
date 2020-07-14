import React, { useState } from "react"

import "./img.less"
import { Modal } from "antd"
interface ImageProps {
    src: string
    width?: number | string
    height?: number | string
}

export const Img: React.FC<ImageProps> = ({
    src,
    width,
    height
}) => {
    const [visible, setVisible] = useState(false)

    return (<div className={"img-box"}
        style={{
            height: height ? height : 214,
            width: width ? width : 160
        }}>
        <Modal
            title="查看图片"
            visible={visible}
            destroyOnClose={true}
            maskClosable={true}
            onCancel={() => setVisible(false)}
            footer={null}
        >
            {src ? <img src={src} alt="图片不存在" /> : "暂无图片"}
        </Modal>
        {src ? <img src={src} alt="图片不存在" onClick={() => setVisible(true)} /> : "暂无图片"}
    </div>)
}
