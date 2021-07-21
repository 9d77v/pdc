import { Spin } from 'antd'
import ReactMarkdown from 'react-markdown'
import { useRecoilValue } from 'recoil'
import 'src/styles/editor.less'
import noteStore from 'src/module/note/note.store'
import gfm from 'remark-gfm'
import Tex from '@matejmazur/react-katex'
import math from 'remark-math'
import 'katex/dist/katex.min.css'
import { FC, Suspense,lazy } from 'react'
const CodeBlock = lazy(() => import('src/components/CodeBlock'))

interface INotePage {
    hideTitle?: boolean
}

const NotePage: FC<INotePage> = ({
    hideTitle = false
}) => {
    const currentNote = useRecoilValue(noteStore.currentNote)
    return (
        <div
            style={{
                width: "100%",
                maxWidth: 760,
                minHeight: 766,
                height: "100%",
                backgroundColor: "#fff",
                wordWrap: "break-word"
            }}>
            {hideTitle ? null : <div style={{ fontSize: 36, height: 56, marginTop: 12, marginBottom: 12, textAlign: 'center', fontWeight: 600, whiteSpace: 'normal' }}>{currentNote.title}</div>}
            <div style={{ paddingLeft: 32, paddingRight: 32, width: "100%", textAlign: 'left' }} >
                <Suspense fallback={<Spin />}>
                    <ReactMarkdown
                        children={currentNote.content || ''}
                        plugins={[[gfm], [math]]}
                        renderers={{
                            inlineMath: ({ value }) => <Tex math={value} />,
                            math: ({ value }) => <Tex block math={value} />,
                            code: CodeBlock
                        }}
                        escapeHtml={false} />
                </Suspense>
            </div>
        </div >

    )
}

export default NotePage
