<template>
  <div class="main" id="main" @contextmenu.prevent="contextMenuHandler">
    <!-- tab目录结构 -->
    <div class="breadCard">
      <div v-if="menuTabs.length > 1"><el-link type="primary">返回上一级</el-link><span class="icon-right">|</span></div>
      <div v-for="(i, index) in menuTabs" :key="index">
        <span>
          <el-link type="primary" v-if="index < menuTabs.length - 1">{{i.name}}</el-link>
          <span v-else :title="i.name" style="color: #666666;" @click="chakan" >{{i.name}}</span>
        </span>
        <span class="icon-right" v-if="index < menuTabs.length - 1">></span>
      </div>
    </div>
    <div class="file-menu" v-if="isMenu">
      <div class="checkTitle">
        <el-checkbox v-model="checkAll">
          <span v-if="multipleSelection.length !== 0">已选中{{multipleSelection.length}}个文件/文件夹</span> 
          <span>全选</span>
         </el-checkbox>
      </div>
      <div class="file-box" v-for="(i,index) in fileData" :key="index" :class="{inputChecked: i.checked}"
        @contextmenu.prevent="fileMenus(i)">
        <!-- 新加文件夹 -->
        <template v-if="i.isCreadte">
          <div class="fileimg"><img :src="menuImg(i.img)" alt=""></div>
          <input type="text" class="addInput" :value="i.val">
          <span class="addBtn"><i class="el-icon-check"></i></span>
          <span class="addBtn" @click="cancelAdd"><i class="el-icon-close"></i></span>
        </template>
        <template v-else>
          <input type="checkbox" @change.stop="checkedStyle(i, $event)" :checked="i.checked" class="el-icon-success">
          <div class="fileimg" @click="openFile(i)"><img :src="menuImg(i.img)" alt=""></div>
          <div class="filename" @click="openFile(i)">{{i.fileName}}</div>
        </template>
      </div>
    </div>
    <div class="file-list" v-else>
       <el-table
          ref="multipleTable"
          :data="fileData"
          tooltip-effect="dark"
          style="width: 100%"
          @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55"></el-table-column>
          <el-table-column prop="fileName" :label="multipleSelection.length === 0 ? '文件名' : '已选中' + multipleSelection.length + '个文件/文件夹'" width="450">
            <template slot-scope="scope" >
              <template v-if="scope.row.isCreadte">
                <img class="listImg" :src="listImg(scope.row.img)" alt="">
                <input type="text" class="addInput" :value="scope.row.val">
                <span class="addBtn"><i class="el-icon-check"></i></span>
                <span class="addBtn" @click="cancelAdd"><i class="el-icon-close"></i></span>
              </template>
              <template v-else>
                <img class="listImg" :src="listImg(scope.row.img)" alt="">
                <span class="filename">{{scope.row.fileName}}</span>
              </template>
            </template>
          </el-table-column>
          <el-table-column label="" prop="fileSize"></el-table-column>
          <el-table-column label="" prop="date"></el-table-column>
        </el-table>
    </div>
    <!-- 鼠标右键菜单列表 -->
    <rightMenu :visible.sync="show" @addFile="addFile"/>
  </div>
</template>
<script>
  import rightMenu from '@/components/rightMenu'
  export default {
    components: {
      rightMenu
    },
    props: {
      isMenu: {
        type: Boolean,
        default: true
      },
      fileData: {
        type: Array,
        default: function () {
          return []
        }
      },
      // 是否创建新文件
      isCreadte: {
        type: Boolean,
        default: true
      }
    },
    data () {
      return {
        target: document.getElementsByClassName('right-menu'),
        triggerShowFn: function triggerShowFn() {},
        triggerHideFn: function triggerHideFn() {},
        x: null,
        y: null,
        style: {},
        show: false,
        multipleSelection: [],
        menuTabs: [
          {
            id: '1',
            name: '全部文件',
            path: '/huodong'
          },
          {
            id: '2',
            name: 'Python',
            path: '/Python'
          },
          {
            id: '3',
            name: 'MongoDB',
            path: '/MongoDB'
          }
        ],
        checkAll: false,
        fileD: 'fileD'
      }
    },
    mounted () {
    },
    watch: {
      show (_show) {
        if (_show) {
          this.bindHideEvents()
        } else {
          this.unbindHideEvents()
        }
      }
    },
    methods: {
      // 右键事件事件处理
      contextMenuHandler (e) {
        this.x = e.clientX
        this.y = e.clientY
        this.target[0].style.top = e.clientY + 'px'
        this.target[0].style.left = e.clientX + 'px'
        this.show = true
      },
      // 绑定隐藏菜单事件
      bindHideEvents () {
        // this.triggerHideFn = this.clickDocumentHandler.bind(this);
        // document.addEventListener('mousedown', this.triggerHideFn)
        // document.addEventListener('mousewheel', this.triggerHideFn)
      },
      // 取消绑定隐藏菜单事件
      unbindHideEvents () {
        document.removeEventListener('mousedown', this.triggerHideFn)
        document.removeEventListener('mousewheel', this.triggerHideFn)
      },
      // 鼠标按压事件处理器
      clickDocumentHandler () {
        this.show = false
      },
      // 加载菜单不同的图片
      menuImg (file) {
        if (file === 'PPt') {
          return require('../assets/PPTbig.png')
        } else if (file === 'Code') {
          return require('../assets/Codebig.png')
        } else if (file === 'ZIP') {
          return require('../assets/ZIPbig.png')
        } else {
          return require('../assets/Folderbig.png')
        }
      },
      listImg (file) {
        if (file === 'PPt') {
          return require('../assets/PPT.png')
        } else if (file === 'Code') {
          return require('../assets/Code.png')
        } else if (file === 'ZIP') {
          return require('../assets/ZIP.png')
        } else {
          return require('../assets/Folder.png')
        }
      },
      handleSelectionChange (val) {
        this.multipleSelection = val
        this.$emit('selectData', val)
      },
      // 取消新增文件
      cancelAdd () {
        this.$emit('cancelAdd')
      },
      addFile () {
        this.$emit('addFile')
      },
      chakan () {
        console.log('fileData', this.fileData)
      },
      // 选中时加载box边框样式
      checkedStyle (data, e) {
        data.checked = e.target.checked
        // 将选中数据存储
        if (data.checked) {
          this.multipleSelection.push(data)
        } else {
          this.multipleSelection.map((_, index)=> {
            if (_.id === data.id) {
              this.multipleSelection.splice(index, 1)
            }
          })
        }
        if (this.multipleSelection.length !== 0) this.$emit('selectData', this.multipleSelection)
      },
      // 打开文件夹
      openFile (data) {
        console.log('data', data)
      },
      // 文件上点击鼠标右键
      fileMenus (data) {
        console.log('fileMenus', data)
      }
    }
  }
</script>
<style>
  .main {
    width: 100%;
    height: calc(100% - 50px);
    font-size: 13px;
  }
  .main .breadCard {
    line-height: 16px;
    height: 18px;
    padding-left: 10px;
  }
  .main .breadCard > div {
    display: inline-block;
  }
  .main .breadCard .icon-right {
    width: 20px;
    display: inline-block;
    text-align: center;
    color: #c5d8f3;
  }
  .file-list, .file-menu{
    width: 100%;
    height: 100%;
  }
  .file-menu .file-box {
    width: 120px;
    height: 127px;
    float: left;
    margin: 2px 0 0 6px;
    text-align: center;
    border: 1px solid #fff;
    position: relative;
    cursor: pointer;
    color: #424e67;
  }
  .file-menu .checkTitle {
    border-top-color: #fff;
    height: 36px;
    line-height: 36px;
    color: #888;
    overflow: hidden;
    border-bottom: 1px solid #f2f6fd;
    padding: 0 0 0 14px;
    border-bottom: 1px solid #f2f6fd;
  }
  .file-menu .addInput {
    width: 58%;
  }
  .file-menu .addBtn {
    margin-left: 0;
  }
  .file-menu .addBtn:last-child {
    margin-left: 3px;
  }
  .file-box:hover {
    border: 1px solid #f1f5fa;
    border-radius: 5px;
    background: #f1f5fa;
  }
  .file-box .fileimg {
    position: relative;
    margin: 24px auto 0;
    width: 84px;
    height: 66px;
    background-repeat: no-repeat;
    overflow: hidden;
  }
  .file-box .filename {
    display: block;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    margin: 6px 5px 5px;
    font-size: 14px;
  }
  .file-box input[type="checkbox"] {
    position: absolute;
    top: 4px;
    left: 4px;
    width: 18px;
    height: 18px;
    appearance: none; /*清除复选框默认样式*/
    cursor: pointer;
    display: none; /**默认不显示 */
  }
  .file-box:hover input[type="checkbox"] {
    display: block;
  }
  .file-box input[type="checkbox"]::before {
    font-size: 18px;
    color: #9FDFFF;
  }
  .file-box input[type="checkbox"]:hover::before {
    color: #7FD5FF;
  }
  /* 选中样式 */
  .inputChecked {
    background-color: #f1f5fa;
    border: 1px solid #90d8ff!important;
    border-radius: 5px;
  }
  .file-box input[type="checkbox"]:checked{
    display: block;
  }
  .file-box input[type="checkbox"]:checked::before {
    color: #00ACFF;
  }
  .listImg {
    vertical-align: top;
    margin-right: 14px;
  }
  .filename {
    white-space: nowrap;
    cursor: pointer;
  }
  .file-list .filename:hover, .file-menu .filename:hover {
    color: #09AAFF;
  }
  .addInput {
    padding: 0 0 0 5px;
    width: 60%;
    height: 24px;
    border: 1px solid #C3EAFF;
    background: #fff;
    border-radius: 2px;
    display: inline-block;
    vertical-align: middle;
    margin: 2px 6px 2px 0;
    color: #666;
  }
  .addBtn {
    color: #09AAFF;
    width: 20px;
    height: 20px;
    display: inline-block;
    border: 1px solid #09AAFF;
    text-align: center;
    position: relative;
    vertical-align: middle;
    margin-left: 4px;
    border-radius: 4px;
    cursor: pointer;
  }
  .addBtn i {
    position: absolute;
    top: 3px;
    left: 3px;
    font-size: 12px;
    font-weight: bolder;
  }
  .addInput:focus {
    outline: 0;
    box-shadow: 0 0 5px 0 #09AAFF;
  }

</style>