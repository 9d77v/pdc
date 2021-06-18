import { Icon, NavBar } from "antd-mobile";
import { useHistory } from "react-router-dom";
import Calculator from "src/components/calculator";

const CalculatorMobile = () => {
    const history = useHistory()
    return (
        <div style={{ height: "100%", textAlign: "center", display: "flex" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                onLeftClick={() => history.goBack()}
            >计算器</NavBar>
            <Calculator marginTop={45} />
        </div>
    );
}

export default CalculatorMobile
