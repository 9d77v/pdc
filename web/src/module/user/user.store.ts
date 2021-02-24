import {
    atom,
    selector,
} from 'recoil';
import { IUser } from './user.model';

class UserStore {
    currentUserInfo = atom<IUser>({
        key: 'currentUserInfo',
        default: {
            uid: "",
            name: "",
            password: "",
            avatar: "",
            roleID: 0,
            gender: 0,
            birthDate: 0,
            ip: "",
        },
    })
    loginDisplay = selector({
        key: "loginDisplay",
        get: ({ get }) => {
            const currentUserInfo = get(this.currentUserInfo)
            return currentUserInfo.roleID > 0 ? "block" : "none"
        }
    })
    ownerDisplay = selector({
        key: "ownerDisplay",
        get: ({ get }) => {
            const currentUserInfo = get(this.currentUserInfo)
            return (currentUserInfo.roleID === 1) ? "block" : "none"
        }
    })
    adminDisplay = selector({
        key: "adminDisplay",
        get: ({ get }) => {
            const currentUserInfo = get(this.currentUserInfo)
            return currentUserInfo.roleID > 0 ? "block" : "none"
        }
    })
}

const userStore = new UserStore()
export default userStore
