import {defineStore} from 'pinia'
import {girlsStore} from './Girls'

export const mainStore = defineStore('main', {
    // state属性，用来存储全局的状态
    state: () => {
        return {
            helloworld: 'Hello World',
            count: 0,
            phone: '15139333888'
        }
    },
    // getters属性，用来监视或者说是计算状态的变化，有缓存的功能
    // getters是有缓存的，虽然调用多次，但是值一样就不会被多次调用
    getters: {
        phoneHidden(state): string {
            console.log('phoneHidden被调用了')
            // return state.phone.toString().replace(/^(\d{3})\d{4}(\d{4})$/, '$1****$2')
            return this.phone.toString().replace(/^(\d{3})\d{4}(\d{4})$/, '$1****$2')
        }
    },
    // actions属性，对state里数据变化的业务逻辑，需求不同，编写逻辑不同，用来修改state全局状态数据的。
    actions: {
        // 在actions中写好逻辑，再调用actions
        changeState() {
            this.count++
            this.helloworld = 'Hello World'
        },
        getList() {
            console.log(girlsStore().list)
        }
    }
})
