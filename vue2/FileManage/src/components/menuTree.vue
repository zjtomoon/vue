<template>
  <ul>
    <li v-for="(i, index) in datas" :key="index">
      <span><i :class="i.icon"></i></span>
      <span>{{i.label}}</span>
    </li>
  </ul>
</template>
<script>
  export default {
    name: 'menuTree',
    props: {
      value: {
        type: [Array, String],
        default: null
      },
      disabled: {
        type: Boolean,
        default: false
      },
      defaultProps: {
        type: Object,
        default: function () {
          return {
            children: 'children',
            label: 'label'
          }
        }
      },
      options: {
        type: Object,
        default: function () {
          return {}
        }
      }
    },
    data () {
      return {
        datas: [
          {
            label: '全部文件',
            path: '/home',
            id: '1',
            icon: 'el-icon-document-copy'
          },
          {
            label: '程四姐视频剪辑',
            path: '/home',
            id: '2'
          },
          {
            label: '文件夹',
            path: '/file',
            id: '3'
          },
          {
            label: '孙敏校全集',
            path: '/home',
            id: '4'
          },
        ]
      }
    },
    created () {
      // this.selected = this.options
    },
    methods: {
      handleNodeClick (data) {
        console.log(data)
      }
    },
    watch: {
      value (val) {
        this.selected = val
      },
      selected (val) {
        this.$emit('input', val)
      }
    }
  }
</script>
<style scoped>
  li {
    color: rgb(102, 102, 102);
    padding-left: 20px;
    cursor: pointer;
    line-height: 38px;
  }
  li span:first-child {
    display: inline-block;
    width: 40px;
    text-align: center;
    font-size: 16px;
  }
  li:hover {
    background: #DCE3E6;
  }
  li a .active {
    color: #3b8cff;
  }
</style>
