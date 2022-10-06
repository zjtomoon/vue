<template>
    <div>
        <nav-header @add="add"></nav-header>
        <nav-main :list="list" @del="del"></nav-main>
        <nav-footer :list="list" @clear="clear"></nav-footer>
        <!--<button @click="send"></button>-->
    </div>
</template>

<script>
    import NavHeader from '@/components/NavHeader/NavHeader.vue'
    import NavMain from '@/components/NavMain/NavMain.vue'
    import NavFooter from '@/components/NavFooter/NavFooter.vue'
    import {defineComponent, computed, ref, onMounted, onUnmounted} from 'vue'
    import {useStore} from 'vuex'
    import {useRouter, useRoute} from 'vue-router'

    export default defineComponent({
        name: 'Home',
        components: {
            NavHeader,
            NavMain,
            NavFooter
        },
        props: {
            msg: {
                type: String,
                required: false,
                default: '默认值'
            }
        },
        setup(props, ctx) {
            // onMounted在setup之后执行
            // onMounted(()=> {
            //     // 组件挂载的过程
            //     // 数据 dom ..
            //     // 发请求
            //     // 数据初始化操作
            //
            // })
            // onUnmounted(()=>{
            //     // 组件卸载时的生命周期
            //     // 清除定时器
            //     // 清除闭包函数
            //
            // })
            // // router是全局的路由对象
            // let router = useRouter()
            // // route是当前的路由对象
            // let route = useRoute()
            // let goto = () => {
            //     // push 函数里可以直接传入跳转的路径
            //     // push 如果是传的对象的形式，就可以传递参数
            //     // back: 回退到上一页
            //     // forward: 去到下一页
            //     // go(正数)  正数代表前进，负数代表后退
            //     router.push({
            //         name:'Home',
            //         path:'/home',
            //         // query传递过去的参数都变成字符串类型
            //         // query传参name path，传递的参数会在地址栏中显示，并且刷新后还在
            //         // params传参只能传name
            //         query:{
            //            name:name.value,
            //            num:num.value,
            //            obj:JSON.stringify(obj)
            //         },
            //         params:{
            //           name:name.value,
            //             num:num.value,
            //             obj:JSON.stringify(obj)
            //         }
            //     })
            // }

            // 子组件传值给父组件
            // let childMsg = ref('我是子组件的数据')
            // let send = () => {
            //     // 通过ctx.emit分发事件
            //     // emit第一个参数是事件名称，第二个是传递的数据
            //     ctx.emit('send',childMsg.value)
            // }
            let store = useStore()
            // console.log("store = " + store)
            let list = computed(() => {
                return store.state.list
            })
            let value = ref('')
            // 添加任务
            let add = (val) => {
                value.value = val
                //先判断有没有存在的任务，如果任务存在不能重复添加
                let flag = true
                list.value.map((item) => {
                    if (item.title === value.value) {
                        flag = false
                        alert('任务已存在')
                    }
                })
                // 调用mutation
                // 没有重复的任务
                if (flag) {
                    store.commit('addTodo', {
                        title: value.value,
                        complete: false
                    })
                    // console.log(val)
                }
            }
            // let num = ref(10)
            // let name = ref('jack')

            // 删除任务
            let del = (val) => {
                // console.log('val = ' + val)
                // 调用删除的mutation
                store.commit('delTodo',val)
            }
            // 清除已完成事件
            let clear = (val) => {
                store.commit('clear',val)
            }
            return {
                add,
                value,
                list,
                del,
                clear
                // childMsg,
                // send
                // num,
                // name
            }
        }
    })
</script>

<style lang="scss" scoped>

</style>