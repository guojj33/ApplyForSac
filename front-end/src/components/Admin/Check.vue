<template>
    <el-card :body-style="{ padding: '0px' }" class="checkCard">
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
      width="110%">
      </el-table-column>
      <el-table-column
      prop="userId"
      label="使用者"
      width="100%">
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
      prop="checkStatus"
      label="签到状态"
      width="100%">
      </el-table-column>
      <el-table-column
      label="签到"
      width="200px">
        <template slot-scope="scope">
            <el-button
            size="mini"
            type="danger"
            @click="CheckAppRecordLate(scope.$index, scope.row)"
            :disabled="CanNotEdit(scope.$index, scope.row)">
            未及时
            </el-button>
            <el-button
            size="mini"
            type="primary"
            @click="CheckAppRecordNormal(scope.$index, scope.row)"
            :disabled="CanNotEdit(scope.$index, scope.row)">
            及时
            </el-button>
        </template>
      </el-table-column>
      </el-table>
    </el-card>
</template>

<script>
import global_ from '../../Global'

export default {
    name: 'Check',
    data() {
        return {
            roomOptions: [{
                    value: 'PianoRoom1',
                    label: '钢琴房1'
                }],
            searchForm: {
                selectedRoom: '',
                selectedDate: '',
            },
            isLoading: false,
            appRecordForms: [],
        };
    },
    created() {
        this.GetAppRecordForms();
    },
    methods: {
        GetAppRecordClassName({row, rowIndex}) {
            console.log(row);
            return 'success-row';
        },
        GetAppRecordForms() {
            this.isLoading = true;
            var createAppRecordForm = function() {
                var appRecordForm = new Object ({
                    appRecordId: '',
                    roomName: '',
                    userId: '',
                    date: '',
                    startTime: '',
                    endTime: '',
                    checkStatus: '',
                });
                return appRecordForm;
            }
            var appRecordForms = [];
            var appRecords = [];
            var self_ = this;
            this.axios.get("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/appRecords")
            .then(function(response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("查询申请成功")
                    appRecords = response.data.AppRecords;
                    console.log(appRecords);
                    for (let index in appRecords) {
                        let ar = appRecords[index];
                        if (ar.ReviewStatus === 0 || ar.ReviewStatus === 2 || ar.ApplyStatus === 1 || ar.ApplyUserId === 'SAC') {    //不显示没有通过或者通过了被取消的申请，SAC 的申请
                            continue;
                        }
                        let appRecordForm = createAppRecordForm();
                        appRecordForm.appRecordId = ar.AppRecordId;
                        appRecordForm.roomName = ar.RoomName;
                        appRecordForm.userId = ar.ApplyUserId;
                        appRecordForm.checkStatus = global_.CheckStatus[ar.CheckStatus];

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
                    alert("查询申请失败")
                }
            })
            .catch(function(error) {
                alert(error);
            });
        },
        CheckAppRecordLate(index, row) {
            let appRecordId = row.appRecordId;
            let updateAppRecordRequest = {
                UpdateField: 'CheckStatus',
                NewValue: '2',
                ValueType: 'int',
            };
            let self_ = this;
            this.axios.put("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/appRecords/" + appRecordId, updateAppRecordRequest)
            .then(function (response) {
                if (response.status === 200) {
                    self_.$alert('设置不准时成功','签到',{
                        confirmButtonText: '确定',
                    });
                    self_.GetAppRecordForms();
                }
            })
            .catch(function (error) {
                alert(error);
            });
        },
        CheckAppRecordNormal(index, row) {
            let appRecordId = row.appRecordId;
            let updateAppRecordRequest = {
                UpdateField: 'CheckStatus',
                NewValue: '1',
                ValueType: 'int',
            };
            let self_ = this;
            this.axios.put("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/appRecords/" + appRecordId, updateAppRecordRequest)
            .then(function (response) {
                if (response.status === 200) {
                    self_.$alert('设置准时成功','签到',{
                        confirmButtonText: '确定',
                    });
                    self_.GetAppRecordForms();
                }
            })
            .catch(function (error) {
                alert(error);
            });
        },
        CanNotEdit(index, row) {
            if (row.checkStatus === "未使用") {
                return false;
            } else {
                return true;
            }
        }
    },
}
</script>

<style scoped>

.checkCard {
    width: 910px;
    padding: 20px;
    margin: auto;
}

</style>