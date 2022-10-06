<template>
    <div class="container">
        <div>
            已完成{{isComplete}}/全部{{ list.length }}
        </div>
        <div v-if="isComplete > 0" class="btn">
            <button @click="clear">清除已完成</button>
        </div>
    </div>
</template>


<script>
    import {defineComponent, ref, computed} from 'vue'

    export default defineComponent({
        name: 'NavFooter',
        props: {
            list: {
                type: Array,
                required: true,
            }
        },
        setup(props, ctx) {
            let isComplete = computed(() => {
                // 过滤已完成的
                let arr = props.list.filter((item) => {
                    return item.complete
                })
                return arr.length
            })
            // let all = ref(3)
            // 清除已完成
            let clear = () => {
                // 过滤未完成的
                let arr = props.list.filter((item) => {
                    return item.complete == false
                })
                ctx.emit('clear',arr)
            }
            return {
                isComplete,
                // all,
                clear
            }
        },
    })
</script>


<style scoped lang='scss'>

</style>

