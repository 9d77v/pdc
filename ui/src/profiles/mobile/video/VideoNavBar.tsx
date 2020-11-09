import { NavBar, Popover, Icon } from "antd-mobile"
import React, { useState } from "react"
import {
    HistoryOutlined
} from '@ant-design/icons'
import Item from "antd-mobile/lib/popover/Item"
import { useHistory } from "react-router-dom"
import { AppPath } from "src/consts/path"

export default function VideoNavBar() {

    const [visible, setVisible] = useState(false)
    const history = useHistory()
    const handleVisibleChange = (visible: boolean) => {
        setVisible(visible)
    }

    const onSelect = (opt: any) => {
        setVisible(false)
        switch (opt.props.children) {
            case "最近播放":
                history.push(AppPath.HISTORY)
                break
        }
    }
    return (
        <NavBar
            mode="light"
            style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
            icon={<Icon type="left" />}
            onLeftClick={() => history.goBack()}
            rightContent={
                <Popover mask
                    visible={visible}
                    overlay={[
                        (<Item key="1" icon={<HistoryOutlined />} >最近播放</Item>),
                    ]}
                    align={{
                        overflow: { adjustY: 0, adjustX: 0 },
                    }}
                    onVisibleChange={handleVisibleChange}
                    onSelect={onSelect}
                >
                    <div style={{
                        height: '100%',
                        padding: '0 15px',
                        marginRight: '-15px',
                        display: 'flex',
                        alignItems: 'center',
                    }}
                    >
                        <Icon type="ellipsis" />
                    </div>
                </Popover>
            }
        >视频 </NavBar>)
}
