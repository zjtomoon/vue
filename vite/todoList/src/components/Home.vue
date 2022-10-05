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
import {defineComponent,ref} from 'vue'

export default defineComponent({
    name:'Home',
    components:{
        NavHeader,
        NavMain,
        NavFooter
    },
    setup(props,ctx){
        let num = ref(10)
        let name = ref('jack')
        return {
            num,
            name
        }
    }
})
</script>

<style lang="scss" scoped>

</style>