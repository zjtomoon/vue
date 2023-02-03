<template>
  <div class="my-multi-table">
    <div class="print" ref="print">
      <div style="margin: 0 auto">

        <div style="text-align: center">中金健康科技（嘉兴）有限公司管理体系记录</div>
        <div>
          <div
              style="color: #666666;text-align: center;margin-top: 30px; margin-bottom: 10px;display: flex;justify-content: space-between"
          >
            <div style="text-align: left">编号:</div>
            <div style="text-align: right">
              <img src="public/static/pic1.png">
            </div>
          </div>
          <div style="color: #666666;text-align: center;margin-top: 30px;margin-bottom: 10px;display: flex;justify-content: space-between">
            <div style="text-align: left">版本号:</div>
          </div>
        </div>
        <h3 style="text-align: center;margin-top: 30px;margin-bottom: 10px">入库预约单</h3>
        <div
            style="color: #666666;text-align: center;margin-top: 30px;margin-bottom: 10px;display: flex;justify-content: space-between"
        >
          <div>顾客名称:</div>
        </div>
        <div style="color: #666666;text-align: center;margin-top: 30px;margin-bottom: 10px;display: flex;justify-content: space-between">
          <div>单号No:</div>
          <div>制单日期:2022/9/3</div>
        </div>
        <table
            style="width: 100%"
            class="table table-striped table-bordered"
            align="center"
            valign="center"
        >
          <thead>
            <tr>
              <td class="column" colspan="8" style="font-weight: bold">单据信息</td>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td class="column" rowspan="2">日期</td>
              <td class="column">预约入场</td>
              <td class="value"></td>
              <td class="column" rowspan="2">地址</td>
              <td class="column">发货方</td>
              <td class="value"></td>
            </tr>
            <tr>
              <td class="column">期望提货</td>
              <td class="value"></td>
              <td class="column">收获方</td>
              <td class="value"></td>
            </tr>
          </tbody>
          <tr>
            <td class="value" colspan="8"></td>
          </tr>
        </table>
        <table style="width: 100%"
               class="table table-striped table-bordered"
               align="center"
               valign="center"
        >
          <thead>
            <tr>
              <td class="column" colspan="100%" style="font-weight: bold">货物清单</td>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td class="column" rowspan="3" colspan="1">产品名称</td>
              <td class="column" rowspan="3" colspan="1">产品订单</td>
              <td class="column" rowspan="3" colspan="1">产品批号</td>
              <td class="column" colspan="6">尺寸</td>
              <td class="column" rowspan="3" colspan="1">重量</td>
              <td class="column" rowspan="3" colspan="1">数量</td>
              <td class="column" rowspan="3" colspan="1">剂量要求(KG)</td>
            </tr>
            <tr>
              <td class="column" colspan="2">长</td>
              <td class="column" colspan="2">宽</td>
              <td class="column" colspan="2">高</td>
            </tr>
            <tr>
              <td class="column">标准</td>
              <td class="column">允差</td>
              <td class="column">标准</td>
              <td class="column">允差</td>
              <td class="column">标准</td>
              <td class="column">允差</td>
            </tr>
            <tr v-for="i in 12" :key="i">
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
              <td class="value" colspan="1"></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "table04"
}
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