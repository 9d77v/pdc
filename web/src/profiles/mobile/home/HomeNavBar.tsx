import { NavBar, Popover, Icon } from "antd-mobile"
import { FC, useState } from "react"
import {
    ScanOutlined
} from '@ant-design/icons'
import Item from "antd-mobile/lib/popover/Item"
import { useHistory } from "react-router-dom"
import { AppPath } from "src/consts/path"

interface IHomeNavBarProps {
    hidden: boolean
}
const HomeNavBar: FC<IHomeNavBarProps> = ({
    hidden
}) => {
    const [visible, setVisible] = useState(false)
    const history = useHistory()
    const handleVisibleChange = (visible: boolean) => {
        setVisible(visible)
    }

    const onSelect = (opt: any) => {
        setVisible(false)
        switch (opt.props.children) {
            case "扫一扫":
                history.push(AppPath.SCANNER)
                break
        }
    }
    return (<NavBar
        mode="light"
        hidden={hidden}
        rightContent={
            <Popover mask
                visible={visible}
                overlay={[
                    (<Item key="1" icon={<ScanOutlined />} >扫一扫</Item>),
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
    >首页 </NavBar>)
}

export default HomeNavBar
