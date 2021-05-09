import React from "react";

export interface ILongPressActionProps {
    item: any
    timeout: number
    onLongPress: () => void
}

export default class LongPressAction extends React.Component<ILongPressActionProps> {

    private buttonPressTimer: any

    constructor(props: ILongPressActionProps, context?: any) {
        super(props, context);
        this.handleButtonPress = this.handleButtonPress.bind(this)
        this.handleButtonRelease = this.handleButtonRelease.bind(this)
    }

    public handleButtonPress() {
        this.buttonPressTimer = setTimeout(() => this.props.onLongPress(), this.props.timeout);
    }

    public handleButtonRelease() {
        clearTimeout(this.buttonPressTimer);
    }

    public render() {
        const { item } = this.props
        return (
            <div onTouchStart={this.handleButtonPress}
                onTouchEnd={this.handleButtonRelease}
                onMouseDown={this.handleButtonPress}
                onMouseUp={this.handleButtonRelease}
                style={{ width: '100%', height: '100%' }}
            >
                {item}
            </div>
        );
    }
}
