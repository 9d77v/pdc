import { atom } from "recoil"

class GlobalStore {
    menuCollapsed = atom<boolean>({
        key: 'menuCollapsed',
        default: false
    })
}

const globalStore = new GlobalStore()
export default globalStore
