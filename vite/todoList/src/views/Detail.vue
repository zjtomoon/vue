<template>
    <div>
        <h1>这是detail组件</h1>
        <child :msg="msg" @send="send"></child>
    </div>
</template>

<script>
    import {defineComponent,ref} from 'vue'
    // import home from '../views/Home.vue'
    export default defineComponent({
        name:'Detail',
        components:{
            home
        },
        setup() {
            let msg = ref('这是父组件的数据')
            let send = (val) => {
                console.log(val)
            }
            return {
                msg,
                send
            }
        }
    })
</script>

<style scoped>

</style>