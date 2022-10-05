<template>
    <div v-for="(item,index) in list" :key="index">
        <div>
            <input type="checkbox" v-model="item.complete">
            {{item.title}}
            <button>删除</button>
        </div>
    </div>
</template>


<script>
    import {defineComponent, ref} from 'vue'

    export default defineComponent({
        name: 'NavMain',
        setup() {
            let list = ref([
                {
                    title:'吃饭',
                    complete:false
                },
                {
                    title:'睡觉',
                    complete:false
                },
                {
                    title:'写代码',
                    complete:true
                },
            ])
            return { //必须加大括号
                list
            }
        },
    })
</script>


<style scoped lang='scss'>

</style>

