<template>
  <div class="my-multi-table">
    <div class="button_group">
      <el-button type="primary" @click="editSend">编辑</el-button>
      <el-button type="primary" @click="printSendTable">打印运输交接单</el-button>
      <el-button type="primary" @click="previewSendTable">查看出库单</el-button>
    </div>
    <div class="print" ref="print">
      <div style="margin:0 auto;">
        <div
            style="color: #666666; text-align:center;margin-top:30px; margin-bottom: 10px; display: flex; justify-content: space-between;"
        >
          <div>XXX有限公司</div>
          <div>2-SMP-MC-01-015</div>
        </div>
        <h3 style="text-align:center;margin-top:30px; margin-bottom: 10px">冷链发货运输交接单</h3>
        <div style="text-align:right; margin-bottom: 10px">编号：_________________</div>
        <table
            style="width: 100%"
            class="table table-striped table-bordered"
            align="center"
            valign="center"
        >
          <tr>
            <td class="column" colspan="8">基本信息（发货人填写）</td>
          </tr>
          <tr>
            <td class="column" colspan="2">供货单位/生产商</td>
            <td class="value" colspan="6"></td>
          </tr>
          <tr>
            <td class="column" colspan="2">收货单位</td>
            <td class="value" colspan="3"></td>
            <td class="column">收件人</td>
            <td class="value" colspan="2"></td>
          </tr>
          <tr>
            <td class="column" colspan="2">收货地址</td>
            <td class="value" colspan="3"></td>
            <td class="column">电话</td>
            <td class="value" colspan="3"></td>
          </tr>
          <tr>
            <td class="column" colspan="2">客户PO号码</td>
            <td class="value" colspan="2"></td>
            <td class="column" colspan="1">发货仓库</td>
            <td class="value" colspan="3"></td>
          </tr>
          <tr>
            <td class="column">序号</td>
            <td class="column">药品名称</td>
            <td class="column">剂型</td>
            <td class="column">规格</td>
            <td class="column">数量</td>
            <td class="column">单位</td>
            <td class="column">批号</td>
            <td class="column">有效期至</td>
          </tr>
          <tr v-for="i in 3" :key="i">
            <td class="value" colspan="1"></td>
            <td class="value" colspan="1"></td>
            <td class="value" colspan="1"></td>
            <td class="value" colspan="1"></td>
            <td class="value" colspan="1"></td>
            <td class="value" colspan="1"></td>
            <td class="value" colspan="1"></td>
            <td class="value" colspan="1"></td>
          </tr>
          <tr>
            <td class="value" colspan="8">
              <div style="display: flex; flex-direction: column;">
                <div style="display: flex; justify-content: flex-start;">审核人意见：</div>
                <div
                    style="display: flex; justify-content: flex-end; margin-top: 20px"
                >审核人日期/时间：__________________</div>
              </div>
            </td>
          </tr>
          <tr>
            <td class="column" colspan="8">发货信息（发货人填写）</td>
          </tr>
          <tr>
            <td class="column" colspan="2">承运公司</td>
            <td class="value" colspan="2"></td>
            <td class="column" colspan="2">启运时间</td>
            <td class="value" colspan="2"></td>
          </tr>
          <tr>
            <td class="column" colspan="2">温度仪器编号</td>
            <td class="value" colspan="2"></td>
            <td class="column" colspan="2">启运温度</td>
            <td class="value" colspan="2"></td>
          </tr>
          <tr>
            <td class="column" colspan="2">保温箱编号</td>
            <td class="value" colspan="2"></td>
            <td class="column" colspan="2">揽件车牌号</td>
            <td class="value" colspan="2"></td>
          </tr>
          <tr>
            <td class="column" colspan="2">客户PO号码</td>
            <td class="value" colspan="2"></td>
            <td class="column" colspan="2">冷链运输单号</td>
            <td class="value" colspan="2"></td>
          </tr>
          <tr>
            <td class="column" colspan="2">发货人/日期</td>
            <td class="value" colspan="2"></td>
            <td class="column" colspan="2">揽件人/日期</td>
            <td class="value" colspan="2"></td>
          </tr>
          <tr>
            <td class="value" colspan="8" style="text-align: left">备注：</td>
          </tr>
          <tr>
            <td class="column" colspan="8">收货信息（收货人填写）</td>
          </tr>
          <tr>
            <td class="column" colspan="2">到达签收时间</td>
            <td class="value" colspan="2">____年__月__日__时__分</td>
            <td class="column" colspan="2">到达时温度</td>
            <td class="value" colspan="2">℃</td>
          </tr>
          <tr>
            <td class="column" colspan="2">运输是否超温</td>
            <td class="value" colspan="2">__是&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;__否</td>
            <td class="column" colspan="2">数量是否正确</td>
            <td class="value" colspan="2">__是&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;__否</td>
          </tr>
          <tr>
            <td class="column" colspan="2">货物是否完好</td>
            <td class="value" colspan="2">__是&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;__否</td>
            <td class="column" colspan="2">物品是否接收</td>
            <td class="value" colspan="2">__是&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;__否</td>
          </tr>
          <tr>
            <td class="column" colspan="2">接收人/日期</td>
            <td class="value" colspan="2"></td>
            <td class="column" colspan="2">配送人/日期</td>
            <td class="value" colspan="2"></td>
          </tr>
          <tr>
            <td class="value" colspan="8" style="text-align: left">备注：</td>
          </tr>
        </table>
        <br />
        <div style="line-height: 30px; color: #333333;">
          <div>*此记录基本信息栏为电子录入记录。</div>
          <div>注1.此随货同行单签收后由承运公司配送人拍照，须汇同此票药物运输过程温度记录数据形成电子档，于签收后24小时内传回给发货委托人，并于一个月内将纸档寄回给发货委托人.</div>
          <div>&nbsp;&nbsp;&nbsp;2.留在物控部保存的单据可无须填写收货信息项.</div>
        </div>
        <br />
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {};
  },
  props: [""],
  watch: {},
  computed: {},
  created() {},
  methods: {
    // 编辑
    editSend() {
      this.$router.push({
        path: "/auditSendDetail",
        query: { type: "sendDetail" }
      });
    },
    // 打印交接单
    printSendTable() {
      this.$print(this.$refs.print);
    },
    // 查看出货单
    previewSendTable() {
      this.$router.push({
        path: "/orderTable",
        query: { type: "orderTable" }
      });
    }
  }
};
</script>

<style scoped>
.button_group {
  margin-top: 30px;
  margin-right: 30px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
.print {
  width: 90%;
  margin: 0 auto;
  max-width: 1280px;
}
.table {
  border-collapse: collapse;
  border-spacing: 0;
  background-color: transparent;
  display: table;
  width: 100%;
  max-width: 100%;
  width: 800px;
  margin: 0 auto;
}
.table td {
  text-align: center;
  vertical-align: middle;
  font-size: 14px;
  font-family: "Arial Normal", "Arial";
  color: #333333;
  padding: 8px 12px;
}
.table-bordered {
  border: 1px solid #ddd;
}
.column {
  width: 30px;
  height: 30px;
  border: 1px solid #333;
  background: #f1f1f1;
}
.value {
  width: 70px;
  height: 30px;
  border: 1px solid #333;
}
</style>
