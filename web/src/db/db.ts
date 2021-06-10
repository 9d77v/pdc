import { nSQL } from "@nano-sql/core";

export async function noteDBInit(): Promise<void> {
    try {
        await nSQL().createDatabase({
            id: "pdc", // can be anything that's a string
            mode: "IDB", // save changes to IndexedDB, WebSQL or SnapDB!
            tables: [ // tables can be created as part of createDatabase or created later with create table queries
                {
                    name: "note",
                    model: {
                        "id:uuid": { pk: true },
                        "parent_id:string": {},
                        "uid:int": {},
                        "note_type:int": {},
                        "level:int": {},
                        "title:string": {},
                        "content:string": {},
                        "tags:string[]": {},
                        "state:int": {},
                        "version:int": {},
                        "sha1:string": {},
                        "sync_status:string": {},
                        "color:string": {},
                        "created_at:date": {},
                        "updated_at:date": {},
                    }, indexes: {
                        // "parent_id:string": {},
                        // "uid:string": {}
                    }
                }
            ],
            version: 1, // current schema/database version
            onVersionUpdate: (prevVersion) => { // migrate versions
                return new Promise((res, rej) => {
                    // switch (prevVersion) {
                    //     case 1:
                    //         // migrate v1 to v2
                    //         res(2);
                    //         break;
                    // }
                })
            }
        })
    } catch (error) {
        // console.log(error)
    }
}
