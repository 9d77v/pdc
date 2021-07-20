import { useHistory } from "react-router-dom"
import { apolloClient } from "src/utils/apollo_client"
import { QrcodeOutlined, LogoutOutlined, LockOutlined, LineChartOutlined ,UserOutlined} from '@ant-design/icons'
import { List } from "antd-mobile"
import { UserBrief } from "src/profiles/mobile/common/UserBrief"
import { AppPath } from "src/consts/path"
const Item = List.Item

export default function MeIndex() {
    const history = useHistory()
    const logout = () => {
        localStorage.clear()
        apolloClient.resetStore()
        history.push(AppPath.LOGIN)
    }
    return (
        <div style={{
            height: "100%",
            display: "flex",
            flexDirection: "column", paddingTop: 20,
            backgroundColor: "#eee"
        }}>
            <UserBrief host={document.location.host} />
            <List renderHeader={() => ''}>
                <Item
                    thumb={<QrcodeOutlined />}
                    arrow="horizontal"
                    onClick={() => history.push(AppPath.UESR_QECODE)}
                >我的二维码名片</Item>
                  <Item
                    thumb={<UserOutlined />}
                    onClick={() => history.push(AppPath.USER_PROFILE)}
                    arrow="horizontal"
                >
                    修改个人信息
                </Item>
                <Item
                    thumb={<LockOutlined />}
                    onClick={() => history.push(AppPath.USER_ACCOUNT)}
                    arrow="horizontal"
                >
                    修改密码
                </Item>
                <Item
                    thumb={<LockOutlined />}
                    onClick={() => history.push(AppPath.USER_GESTURE_PASSWORD)}
                    arrow="horizontal"
                >
                    手势密码
                </Item>
                <Item
                    thumb={<LineChartOutlined />}
                    onClick={() => history.push(AppPath.USER_DATA_ANALYSIS)}
                    arrow="horizontal"
                >
                    数据统计
                </Item>
                <Item
                    thumb={<LogoutOutlined />}
                    onClick={() => logout()}
                    arrow="horizontal"
                >
                    退出登录
                </Item>
            </List>
        </div>)
}
