import {defineStore} from 'pinia'

export const girlsStore = defineStore('girls',{
    state:()=> {
        return {
            list:['小红','小美','二丫']
        }
    }
})