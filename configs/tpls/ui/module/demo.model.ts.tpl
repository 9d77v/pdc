[[.HasTime]]
export interface I[[.Name]] {[[range .Columns]][[if eq .TSType "dayjs.Dayjs"]]
    [[.Name]]?: [[.TSType]],[[else]]
    [[.Name]]: [[.TSType]],[[end]][[end]]
}

export interface IUpdate[[.Name]] {
    id: number,[[range .Columns]][[if eq .TSType "dayjs.Dayjs"]]
    [[.Name]]?: [[.TSType]],[[else]]
    [[.Name]]: [[.TSType]],[[end]][[end]]
}
