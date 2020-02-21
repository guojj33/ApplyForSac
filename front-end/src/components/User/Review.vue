<template>
    <el-card :body-style="{ padding: '0px' }" class="reviewCard">
      <el-table
      v-loading="isLoading"
      :data="appRecordForms"
      stype="width: 100%">
      <el-table-column
      prop="appRecordId"
      label="审批号"
      width="100%">
      </el-table-column>
      <el-table-column
      prop="roomName"
      label="房间名"
      sortable
      width="110%">
      </el-table-column>
      <el-table-column
      prop="date"
      label="日期"
      sortable
      width="100%">
      </el-table-column>
      <el-table-column
      prop="startTime"
      label="开始时间"
      width="100%">
      </el-table-column>
      <el-table-column
      prop="endTime"
      label="结束时间"
      width="100%">
      </el-table-column>
      <el-table-column
      prop="reviewStatus"
      label="审核状态"
      width="100%">
      </el-table-column>
      <el-table-column
      label="编辑"
      width="100%">
        <template slot-scope="scope">
            <el-button
                size="mini"
                type="danger"
                :disabled="CanNotCancel(scope.$index, scope.row)"
                @click="CancelAppRecord(scope.$index, scope.row)">
                取消
            </el-button>
        </template>
      </el-table-column>
      </el-table>
    </el-card>
</template>

<script>
import global_ from '../../Global'

export default {
    name: 'Review',
    data() {
        return {
            isLoading: true,
            appRecordForms: [],
        }
    },
    created() {
        this.GetAppRecordForms();
    },
    destroyed() {
        this.appRecordForms = '';
    },
    methods: {
        GetAppRecordForms() {
            this.isLoading = true;
            var createAppRecordForm = function() {
                var appRecordForm = new Object ({
                    appRecordId: '',
                    roomName: '',
                    date: '',
                    startTime: '',
                    endTime: '',
                    reviewStatus: '',
                    applyStatus: '',
                });
                return appRecordForm;
            }
            let appRecordForms = [];
            let appRecords = [];
            let self_ = this;
            this.axios.get("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/appRecords")
            .then(function(response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("获取申请成功")
                    appRecords = response.data.AppRecords;
                    //console.log(appRecords);
                    for (let index in appRecords) {
                        let ar = appRecords[index];
                        console.log(ar);
                        let appRecordForm = createAppRecordForm();
                        appRecordForm.appRecordId = ar.AppRecordId;
                        appRecordForm.roomName = ar.RoomName;
                        appRecordForm.reviewStatus = global_.ReviewStatus[ar.ReviewStatus];
                        appRecordForm.applyStatus = global_.ApplyStatus[ar.ApplyStatus];

                        let duration = ar.ApplyUsingTime;
                        let startTime = new Date(duration.StartTime);
                        let endTime = new Date(duration.EndTime);
                        appRecordForm.date = self_.$moment(startTime).utc().format("YYYY-MM-DD");
                        appRecordForm.startTime = self_.$moment(startTime).utc().format("HH:mm");
                        appRecordForm.endTime = self_.$moment(endTime).utc().format("HH:mm");
                        

                        appRecordForms.push(appRecordForm);
                    }
                    console.log(appRecordForms)
                    self_.appRecordForms = appRecordForms;
                    self_.isLoading = false;
                } else {
                    alert("获取申请失败")
                }
            })
            .catch(function(error) {
                alert(error);
            });
        },
        CancelAppRecord(index, row) {
            var appRecordId = row.appRecordId;
            var updateAppRecordRequest = {
                UpdateField: 'ApplyStatus',
                NewValue: '1',
                ValueType: 'int',
            };
            var self_ = this;
            this.axios.put("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/appRecords/" + appRecordId, updateAppRecordRequest)
            .then(function (response) {
                if (response.status === 200) {
                    self_.$alert('取消成功','申请',{
                        confirmButtonText: '确定',
                    });
                    self_.GetAppRecordForms();
                }
            })
            .catch(function (error) {
                alert(error);
            });
        },
        CanNotCancel(index, row) {
            console.log(row.applyStatus);
            if (row.applyStatus === "已取消" || row.reviewStatus === "未通过") {
                return true;
            } else {
                return false;
            }
        }
    }
}
</script>

<style scoped>

.reviewCard {
    width: 720px;
    padding: 20px;
    margin: auto;
}

</style>