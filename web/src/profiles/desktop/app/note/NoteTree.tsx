import { CascaderOptionType, CascaderValueType } from 'antd/lib/cascader';
import { FC } from 'react';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil'
import CustomCascader from 'src/components/CustomCascader';
import noteStore from 'src/module/note/note.store'

interface INoteTreeProps {
  updateCurrentNote: (id: string, editable: boolean) => Promise<void>
}

const NoteTree: FC<INoteTreeProps> = ({ updateCurrentNote }) => {
  const noteTrees = useRecoilValue(noteStore.noteTrees)
  const setNotes = useSetRecoilState(noteStore.notes)
  const [noteSelectIDs, setNoteSelectIDs] = useRecoilState(noteStore.noteSelectIDs)
  const currentNote = useRecoilValue(noteStore.currentNote)
  const onChange = async (value: CascaderValueType, selectedOptions?: CascaderOptionType[] | undefined) => {
    let id: any = 'root'
    if (value && value.length > 0) {
      id = value[value.length - 1]
    }
    setNoteSelectIDs(value)
    let editable = currentNote.editable || false
    if (id !== currentNote.id) {
      editable = false
    }
    await updateCurrentNote(id, editable)
    if (value.length < 3) {
      const notes = await noteStore.findByParentID(id, currentNote.uid)
      setNotes(notes)
    }
  }

  return (
    <div style={{ width: 300, float: 'left' }}>
      <CustomCascader
        options={noteTrees}
        value={noteSelectIDs}
        onChange={onChange}
      />
    </div>
  )
}

export default NoteTree
