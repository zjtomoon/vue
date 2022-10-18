import { defineStore } from 'pinia'

export const mainStore = defineStore('main',{
    // state属性，用来存储全局的状态
    state:()=> {
        return {
            helloworld:'Hello World',
            count:0
        }
    },
    // getters属性，用来监视或者说是计算状态的变化，有缓存的功能
    getters:{},
    // actions属性，对state里数据变化的业务逻辑，需求不同，编写逻辑不同，用来修改state全局状态数据的。
    actions:{
        // 在actions中写好逻辑，再调用actions
        changeState(){
            this.count++
            this.helloworld = 'Hello World'
        }
    }
})