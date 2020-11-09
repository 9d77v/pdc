
export function checkIsFile(source: any) {
    return source instanceof File
}

export function getExt(url: string): string {
    if (url.includes('?')) {
        return getExt(url.split('?')[0])
    }

    if (url.includes('#')) {
        return getExt(url.split('#')[0])
    }

    return url
        .trim()
        .toLowerCase()
        .split('.')
        .pop() || ''
}


export function getType(source: any) {
    return checkIsFile(source) ? getExt(source.name) : getExt(source)
}

export function subToVtt(sub: any) {
    return (
        'WEBVTT\n\n' +
        sub
            .map((item: any, index: any) => {
                return index + 1 + '\n' + item.start + ' --> ' + item.end + '\n' + item.text
            })
            .join('\n\n')
    )
}

export async function getVttFromFile(file: File) {
    return new Promise<string>(resolve => {
        const reader = new FileReader()
        reader.onload = () => {
            switch (getType(file)) {
                case 'vtt':
                    resolve(reader.result?.toString().replace(/{[\s\S]*?}/g, ''))
                    break
                case 'ass':
                    resolve(assToVtt(reader.result))
                    break
                case 'srt':
                    resolve(srtToVtt(reader.result))
                    break
                default:
                    resolve()
                    break
            }
        }
        reader.readAsText(file)
    })
}

export function assToVtt(ass: any) {
    const re_ass = new RegExp(
        'Dialogue:\\s\\d,' +
        '(\\d+:\\d\\d:\\d\\d.\\d\\d),' +
        '(\\d+:\\d\\d:\\d\\d.\\d\\d),' +
        '([^,]*),' +
        '([^,]*),' +
        '(?:[^,]*,){4}' +
        '([\\s\\S]*)$',
        'i',
    )

    function fixTime(time = '') {
        return time
            .split(/[:.]/)
            .map((item, index, arr) => {
                if (index === arr.length - 1) {
                    if (item.length === 1) {
                        return '.' + item + '00'
                    } else if (item.length === 2) {
                        return '.' + item + '0'
                    }
                } else {
                    if (item.length === 1) {
                        return (index === 0 ? '0' : ':0') + item
                    }
                }

                return index === 0 ? item : index === arr.length - 1 ? '.' + item : ':' + item
            })
            .join('')
    }

    return (
        'WEBVTT\n\n' +
        ass
            .split(/\r?\n/)
            .map((line: any) => {
                const m = line.match(re_ass)
                if (!m) return null
                return {
                    start: fixTime(m[1].trim()),
                    end: fixTime(m[2].trim()),
                    text: m[5]
                        .replace(/{[\s\S]*?}/g, '')
                        .replace(/(\\N)/g, '\n')
                        .trim()
                        .split(/\r?\n/)
                        .map((item: any) => item.trim())
                        .join('\n'),
                }
            })
            .filter((line: any) => line)
            .map((line: any, index: any) => {
                if (line) {
                    return index + 1 + '\n' + line.start + ' --> ' + line.end + '\n' + line.text
                } else {
                    return ''
                }
            })
            .filter((line: any) => line.trim())
            .join('\n\n')
    )
}

export function srtToVtt(srt: any) {
    return 'WEBVTT \r\n\r\n'.concat(
        srt
            .replace(/\{\\([ibu])\}/g, '</$1>')
            .replace(/\{\\([ibu])1\}/g, '<$1>')
            .replace(/\{([ibu])\}/g, '<$1>')
            .replace(/\{\/([ibu])\}/g, '</$1>')
            .replace(/(\d\d:\d\d:\d\d),(\d\d\d)/g, '$1.$2')
            .replace(/{[\s\S]*?}/g, '')
            .concat('\r\n\r\n'),
    )
}
