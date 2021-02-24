import { FC } from 'react';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { monokai } from 'react-syntax-highlighter/dist/esm/styles/hljs';

export interface ICodeBlockProps {
    value: string,
    language?: string
}

const CodeBlock: FC<ICodeBlockProps> = ({ value, language }) => {
    return (
        <SyntaxHighlighter language={language} style={monokai} showLineNumbers={true} >
            {value || ''}
        </SyntaxHighlighter>
    )
}

export default CodeBlock
