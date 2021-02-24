import { Spin } from 'antd'
import * as React from 'react'
import ReactMarkdown from 'react-markdown'
import { useRecoilValue } from 'recoil'
import 'src/styles/editor.less'
import noteStore from 'src/module/note/note.store'
import CodeBlock from 'src/components/CodeBlock'

const NotePage = () => {
    const currentNote = useRecoilValue(noteStore.currentNote)

    return (
        <div style={{ display: 'inline-flex', alignItems: 'center', justifyContent: 'center', width: '100%', height: '100%', paddingTop: 12 }}>
            <div
                style={{
                    width: "100%",
                    maxWidth: 760,
                    height: "100%",
                    minHeight: 760,
                    backgroundColor: "#fff", boxShadow: '3px 3px 3px 3px darkgrey',
                    marginBottom: 16, marginLeft: 12, marginRight: 12,
                }}>
                <div style={{ fontSize: 36, height: 56, marginTop: 24, marginBottom: 24, textAlign: 'center', fontWeight: 600, whiteSpace: 'normal' }}>{currentNote.title}</div>
                <div style={{ margin: "0 24px", maxWidth: 666, width: "100%", textAlign: 'left' }} >
                    <React.Suspense fallback={<Spin />}>
                        <ReactMarkdown source={currentNote.content || ''} renderers={{ code: CodeBlock }} escapeHtml={false} />
                    </React.Suspense>
                </div>
            </div>
        </div >

    )
}

export default NotePage
