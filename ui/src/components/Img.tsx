import React from "react"

import "./img.less"
interface ImageProps {
    src: string
    width?: number
    height?: number
}

export const Img: React.FC<ImageProps> = ({
    src,
    width,
    height
}) => {

    return (<div className={"img-box"}
        style={{
            height: height ? height : 214,
            width: width ? width : 160
        }}>
        {src ? <img src={src} alt="图片不存在" /> : "暂无图片"}
    </div>)
}
