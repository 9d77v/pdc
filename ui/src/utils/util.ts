import moment from "moment"
function isMobile(): boolean {
    return /Android|webOS|iPhone|iPod|BlackBerry/i.test(navigator.userAgent)
}

function isIPhone(): boolean {
    return /iPhone|iPod/i.test(navigator.userAgent)
}

function formatTime(t: number): string {
    const m = moment(t * 1000)
    if (t < 3600) {
        return m.format('mm:ss')
    }
    return m.format('HH:mm:ss')
}

function formatDetailTime(t: number): string {
    const m = moment(t * 1000)
    return m.format('YYYY-MM-DD HH:mm:ss')
}

function formatRelativeTime(t: number): string {
    const m = moment(t * 1000)
    const day = m.format("YYYY-MM-DD")
    if (day === moment().format("YYYY-MM-DD")) {
        return m.format("HH:mm")
    }
    if (day === moment().add("-1", "days").format("YYYY-MM-DD")) {
        return "昨天 " + m.format("HH:mm")
    }
    return day
}

export { isMobile, formatTime, formatDetailTime, formatRelativeTime, isIPhone }