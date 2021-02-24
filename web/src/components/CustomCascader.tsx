import { Cascader } from 'antd'
import { CascaderOptionType, CascaderValueType } from 'antd/lib/cascader';
import { FC } from 'react';

export interface ICustomCascaderProps {
    options?: CascaderOptionType[] | undefined
    value?: CascaderValueType | undefined
    onChange?: ((value: CascaderValueType, selectedOptions?: CascaderOptionType[] | undefined) => void) | undefined
}

const specialCharater = new RegExp("[`~!@#$%^&*()\\-+={}':;,\\[\\].<>/?￥…（）_|【】‘；：”“’。，、？\\s]");

const displayRender = (labels: any, selectedOptions: CascaderOptionType[] | undefined) =>
    labels.map((label: any, i: number) => {
        if (selectedOptions) {
            const option = selectedOptions[i];
            if (i === labels.length - 1) {
                return <span key={i}>{option?.title} </span>;
            }
            return <span key={i}>{option.title} / </span>;
        }
        return <span key={i} />
    });

const CustomCascader: FC<ICustomCascaderProps> = ({ options, value, onChange }) => {
    const heightLight = (str: string, keyword: string, prefixCls: any) => {
        const keyArr = keyword.split(specialCharater).filter(k => k);
        const newText = str.replace(
            new RegExp(keyArr.join("|"), "ig"),
            substr => `<span class="${"".concat(prefixCls, "-menu-item-keyword")}">${substr}</span>`
        );
        return (<span dangerouslySetInnerHTML={{ __html: newText }} />);
    }

    const defaultFilter = (inputValue: string, path: any, names: any): boolean => {
        return path.some((option: any) => option.title.toLowerCase().indexOf(inputValue.toLowerCase()) > -1);
    }

    const defaultRenderFilteredOption = (inputValue: any, path: any, prefixCls: any, names: any) => {
        return path.map((option: any, index: any) => {
            const label = option.title;
            const node = label.toLowerCase().indexOf(inputValue.toLowerCase()) > -1 ? heightLight(label, inputValue, prefixCls) : label;
            return index === 0 ? node : [' / ', node];
        });
    }

    const defaultSortFilteredOption = (a: any, b: any, inputValue: any, names: any) => {
        function callback(elem: any): any {
            return elem.title.toLowerCase().indexOf(inputValue.toLowerCase()) > -1;
        }

        return a.findIndex(callback) - b.findIndex(callback);
    }

    return (
        <Cascader
            showSearch={{ filter: defaultFilter, render: defaultRenderFilteredOption, sort: defaultSortFilteredOption }}
            options={options}
            expandTrigger="hover"
            value={value}
            onChange={onChange}
            displayRender={displayRender}
            changeOnSelect={true}
            placeholder="搜索"
            style={{ width: '100%' }}
        />
    )
}

export default CustomCascader
