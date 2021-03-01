import { Icon, List, Modal, SwipeAction, Toast } from 'antd-mobile';
import { useRecoilValue } from 'recoil';
import { LongPressAction } from 'src/components/LongPressAction';
import { INote, NoteType } from 'src/module/note/note.model';
import noteStore from 'src/module/note/note.store';

const Item = List.Item;
const alert = Modal.alert
const prompt = Modal.prompt
const operation = Modal.operation


const NoteList = () => {
    const notes = useRecoilValue(noteStore.notes)

    const onClick = (index: number) => {
        // this.props.noteStore.nextNoteList(this.props.noteStore.notes[index])
    }

    const onLongPress = (note: INote) => {
        operation([
            { text: '重命名', onPress: () => onRename(note) },
            { text: '删除', onPress: () => onDelete(note) },
        ])
    }

    const onRename = (note: INote) => {
        prompt(
            '新标题',
            '',
            async (title) => {
                if (title.length > 20) {
                    Toast.info('标题长度不能超过20', 1);
                } else {
                    // await this.props.noteStore.updateNote(note, title)
                    // await this.props.noteStore.listLocalNote(this.props.noteStore.currentNote.id)
                }
            },
            'default',
            note.title,
            ['请输入标题'], /Android/i.test(navigator.userAgent) ? 'android' : 'ios'
        )
    }

    const onDelete = (note: INote) => {
        let confirmText = ""
        if (note.note_type === NoteType.Directory) {
            confirmText = '确认删除笔记本[' + note.title + ']？删除后笔记本下所有的内容都无法访问'
        } else {
            confirmText = "确认删除笔记？"
        }

        alert('确认删除', confirmText, [
            {
                text: '确认', onPress: async () => {
                    // await this.props.noteStore.hideNote(note)
                    // setTimeout(() => {
                    //     this.props.noteStore.listLocalNote(this.props.noteStore.currentNote.id)
                    // }, 300)

                }
            },
            { text: '取消' }
        ],
            /Android/i.test(navigator.userAgent) ? 'android' : 'ios'
        )
    }

    let listItems: any[] = []
    if (notes) {
        listItems = notes.map((item: INote, i: number) => {
            const listItem = <Item
                style={{ height: 60 }}
                onClick={() => onClick(i)}
            ><Icon style={{
                flex: 1, color: item.note_type === NoteType.Directory ? "brown" : "blue",
                paddingRight: 20, paddingLeft: 6, fontSize: 32
            }}
                type={item.note_type === NoteType.Directory ? "folder" : "file"} />
                {item.title === undefined ? '' : item.title}
            </Item>
            if (/Android/i.test(navigator.userAgent)) {
                return (
                    <LongPressAction key={item.id}
                        timeout={500}
                        onLongPress={onLongPress.bind(this, item)} item={listItem}

                    />
                )
            }
            return (
                <SwipeAction key={item.id}
                    style={{ backgroundColor: 'gray' }}
                    autoClose={true}
                    right={[
                        {
                            text: '重命名',
                            onPress: () => onRename(item),
                            style: { backgroundColor: '#ddd', color: 'white', width: 60 },
                        },
                        {
                            text: '删除',
                            onPress: () => onDelete(item),
                            style: { backgroundColor: '#F4333C', color: 'white', width: 60 },
                        },
                    ]}
                >
                    {listItem}
                </SwipeAction>

            )
        }
        )
    }
    return (
        <List >
            {listItems}
        </List>
    )
}


export default NoteList
