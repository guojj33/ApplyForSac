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
      width="100%">
      </el-table-column>
      <el-table-column
      prop="userId"
      label="使用者"
      width="100%">
      </el-table-column>
      <el-table-column
      prop="date"
      label="日期"
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
      prop="descript"
      label="理由"
      width="100%">
      </el-table-column>
      <el-table-column
      prop="reviewStatus"
      label="审核状态"
      width="100%">
      </el-table-column>
      <el-table-column
      label="审核"
      width="150px">
        <template slot-scope="scope">
            <el-button
            size="mini"
            type="danger"
            @click="RejectAppRecord(scope.$index, scope.row)"
            :disabled="CanNotEdit(scope.$index, scope.row)">
            拒绝
            </el-button>
            <el-button
            size="mini"
            type="primary"
            @click="AcceptAppRecord(scope.$index, scope.row)"
            :disabled="CanNotEdit(scope.$index, scope.row)">
            通过
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
                    descript: '',
                    reviewStatus: '',
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
                        if (ar.ApplyStatus === 1) {//已经取消的申请
                            continue;
                        }
                        let appRecordForm = createAppRecordForm();
                        appRecordForm.appRecordId = ar.AppRecordId;
                        appRecordForm.roomName = global_.Eng2ChiRoomName[ar.RoomName];
                        appRecordForm.userId = ar.ApplyUserId;
                        appRecordForm.reviewStatus = global_.ReviewStatus[ar.ReviewStatus];

                        let duration = ar.ApplyUsingTime;
                        let startTime = new Date(duration.StartTime);
                        let endTime = new Date(duration.EndTime);
                        appRecordForm.date = self_.$moment(startTime).format("YYYY-MM-DD");
                        appRecordForm.startTime = self_.$moment(startTime).format("HH:mm");
                        appRecordForm.endTime = self_.$moment(endTime).format("HH:mm");

                        appRecordForm.descript = ar.Description;
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
        RejectAppRecord(index, row) {
            let appRecordId = row.appRecordId;
            let updateAppRecordRequest = {
                UpdateField: 'ReviewStatus',
                NewValue: '2',
                ValueType: 'int',
            };
            let self_ = this;
            this.axios.put("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/appRecords/" + appRecordId, updateAppRecordRequest)
            .then(function (response) {
                if (response.status === 200) {
                    self_.$alert('拒绝成功','审核',{
                        confirmButtonText: '确定',
                    });
                    self_.GetAppRecordForms();
                }
            })
            .catch(function (error) {
                alert(error);
            });
        },
        AcceptAppRecord(index, row) {
            let appRecordId = row.appRecordId;
            let updateAppRecordRequest = {
                UpdateField: 'ReviewStatus',
                NewValue: '1',
                ValueType: 'int',
            };
            let self_ = this;
            this.axios.put("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/appRecords/" + appRecordId, updateAppRecordRequest)
            .then(function (response) {
                if (response.status === 200) {
                    self_.$alert('通过成功','审核',{
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
            //console.log(row.reviewStatus === "待审核");
            if (row.reviewStatus === "待审核") {
                return false;
            } else {
                return true;
            }
        }
    },
}
</script>

<style scoped>

.reviewCard {
    width: 950px;
    padding: 20px;
    margin: auto;
}

</style>