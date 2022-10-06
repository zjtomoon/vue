<template>
    <div>
        <div v-if="list.length > 0">
            <div v-for="(item,index) in list" :key="index">
                <div class="item">
                    <input type="checkbox" v-model="item.complete">
                    {{item.title}}
                    <button class="del" @click="del(item,index)">删除</button>
                </div>
            </div>
        </div>
        <div v-else>
            暂无任务
        </div>
    </div>
</template>


<script>
    import {defineComponent, ref} from 'vue'

    export default defineComponent({
        name: 'NavMain',
        props: {
            list: {
                type: Array,
                required: true,
            }
        },
        emits: ['del'],
        setup(props, ctx) {
            // let list = ref([
            //     {
            //         title:'吃饭',
            //         complete:false
            //     },
            //     {
            //         title:'睡觉',
            //         complete:false
            //     },
            //     {
            //         title:'写代码',
            //         complete:true
            //     },
            // ])
            // 删除任务
            let del = (item, index) => {
                ctx.emit('del', index)
                // console.log(item)
                // console.log(index)
            }
            return { //必须加大括号
                // list,
                del
            }
        },
    })
</script>


<style scoped lang='scss'>
    .item {
        height: 35px;
        line-height: 35px;
        position: relative;
        width: 160px;
        cursor: pointer;
        button {
            position: absolute;
            right: 20px;
            top: 6px;
            display: none;
            z-index: 99;
        }
        &:hover {
            background: #ddd;
            button {
                display: block;
            }
        }
    }
</style>

