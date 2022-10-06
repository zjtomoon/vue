<template>
    <div>
        <nav-header></nav-header>
        <nav-main></nav-main>
        <nav-footer></nav-footer>
    </div>
</template>

<script>
import NavHeader from '@/components/NavHeader/NavHeader.vue'
import NavMain from '@/components/NavMain/NavMain.vue'
import NavFooter from '@/components/NavFooter/NavFooter.vue'
import {defineComponent,computed,ref,onMounted,onUnmounted} from 'vue'
import {useStore} from 'vuex'
import {useRouter,useRoute} from 'vue-router'

export default defineComponent({
    name:'Home',
    components:{
        NavHeader,
        NavMain,
        NavFooter
    },
    setup(props,ctx){
        // onMounted在setup之后执行
        onMounted(()=> {
            // 组件挂载的过程
            // 数据 dom ..
            // 发请求
            // 数据初始化操作

        })
        onUnmounted(()=>{
            // 组件卸载时的生命周期
            // 清除定时器
            // 清除闭包函数

        })
        // router是全局的路由对象
        let router = useRouter()
        // route是当前的路由对象
        let route = useRoute()
        let goto = () => {
            // push 函数里可以直接传入跳转的路径
            // push 如果是传的对象的形式，就可以传递参数
            // back: 回退到上一页
            // forward: 去到下一页
            // go(正数)  正数代表前进，负数代表后退
            router.push({
                name:'Home',
                path:'/home',
                // query传递过去的参数都变成字符串类型
                // query传参name path，传递的参数会在地址栏中显示，并且刷新后还在
                // params传参只能传name
                query:{
                   name:name.value,
                   num:num.value,
                   obj:JSON.stringify(obj)
                },
                params:{
                  name:name.value,
                    num:num.value,
                    obj:JSON.stringify(obj)
                }
            })
        }
        let store = useStore()
        let list = computed(() => {
            return store.state.list
        } )
        // let num = ref(10)
        // let name = ref('jack')
        return {
            // num,
            // name
        }
    }
})
</script>

<style lang="scss" scoped>

</style>