import { Spin } from 'antd'
import * as React from 'react'
import ReactMarkdown from 'react-markdown'
import { useRecoilValue } from 'recoil'
import 'src/styles/editor.less'
import noteStore from 'src/module/note/note.store'
import gfm from 'remark-gfm'
import Tex from '@matejmazur/react-katex'
import math from 'remark-math'
import 'katex/dist/katex.min.css'

const CodeBlock = React.lazy(() => import('src/components/CodeBlock'))
const NotePage = () => {
    const currentNote = useRecoilValue(noteStore.currentNote)

    return (
        <div style={{ display: 'inline-flex', alignItems: 'center', justifyContent: 'center', width: '100%', height: '100%', paddingTop: 12, marginBottom: 32 }}>
            <div
                style={{
                    width: "100%",
                    maxWidth: 760,
                    minHeight: 766,
                    height: "100%",
                    backgroundColor: "#fff", boxShadow: '3px 3px 3px 3px darkgrey',
                    marginLeft: 12, marginRight: 12
                }}>
                <div style={{ fontSize: 36, height: 56, marginTop: 24, marginBottom: 24, textAlign: 'center', fontWeight: 600, whiteSpace: 'normal' }}>{currentNote.title}</div>
                <div style={{ margin: "0 24px", maxWidth: 666, width: "100%", textAlign: 'left' }} >
                    <React.Suspense fallback={<Spin />}>
                        <ReactMarkdown
                            children={currentNote.content || ''}
                            plugins={[[gfm], [math]]}
                            renderers={{
                                inlineMath: ({ value }) => <Tex math={value} />,
                                math: ({ value }) => <Tex block math={value} />,
                                code: CodeBlock
                            }}
                            escapeHtml={false} />
                    </React.Suspense>
                </div>
            </div>
        </div >

    )
}

export default NotePage
