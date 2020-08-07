import moment from "moment"
function isMobile(): boolean {
    return /Android|webOS|iPhone|iPod|BlackBerry/i.test(navigator.userAgent)
}

function formatTime(t: number): string {
    const m = moment(t * 1000)
    if (t < 3600) {
        return m.format('mm:ss')
    }
    return ""
}

function formatRelativeTime(t: number): string {
    const m = moment(t * 1000)
    const day = m.format("YYYY-MM-DD")
    if (day === moment().format("YYYY-MM-DD")) {
        return m.format("hh:mm")
    }
    return day
}

export { isMobile, formatTime, formatRelativeTime }