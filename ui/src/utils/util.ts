import dayjs from 'dayjs'
function isMobile(): boolean {
    return /Android|webOS|iPhone|iPod|BlackBerry/i.test(navigator.userAgent)
}

function isIPhone(): boolean {
    return /iPhone|iPod/i.test(navigator.userAgent)
}

function formatTimeLength(t: number): string {
    const m = dayjs(t * 1000)
    if (t < 3600) {
        return m.format('mm:ss')
    }
    const h = parseInt((t / 3600).toString(), 10)
    if (h < 10) {
        return "0" + h + ":" + m.format('mm:ss')
    }
    return h + ":" + m.format('mm:ss')
}

function formatDetailTime(t: number): string {
    const m = dayjs(t * 1000)
    return m.format('YYYY-MM-DD HH:mm:ss')
}

function formatRelativeTime(t: number): string {
    const m = dayjs(t * 1000)
    const day = m.format("YYYY-MM-DD")
    if (day === dayjs().format("YYYY-MM-DD")) {
        return m.format("HH:mm")
    }
    if (day === dayjs().add(-1, "day").format("YYYY-MM-DD")) {
        return "昨天 " + m.format("HH:mm")
    }
    return day
}

export { isMobile, formatTimeLength, formatDetailTime, formatRelativeTime, isIPhone }