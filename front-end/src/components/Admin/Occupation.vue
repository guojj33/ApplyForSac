<template>
    <el-card :body-style="{ padding: '0px' }" class="occupCard">
        <el-form :model="searchForm">
            <el-form-item>
                <el-select v-model="searchForm.selectedRoom" placeholder="请选择房间" @change="GetAppRecordForms()" id="roomPicker" class="input">
                    <el-option
                        v-for="item in roomOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                        ></el-option>
                </el-select>
                <el-date-picker
                id="datePicker"
                v-model="searchForm.selectedDate"
                type="date"
                placeholder="请选择日期" class="input">
                </el-date-picker>
            </el-form-item>
        </el-form>

      <el-table
      v-loading="isLoading"
      :data="appRecordForms"
      stype="width: 100%"
      :row-class-name="GetAppRecordClassName">
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
      prop="reviewStatus"
      label="审核状态"
      width="100%">
      </el-table-column>
      </el-table>

    </el-card>
</template>

<script>
import global_ from '../../Global'

export default {
    name: 'Occupation',
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
                    date: '',
                    startTime: '',
                    endTime: '',
                    reviewStatus: '',
                });
                return appRecordForm;
            }
            var appRecordForms = [];
            var appRecords = [];
            var self_ = this;
            this.axios.get("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/rooms/" + self_.searchForm.selectedRoom + "/appRecords")
            .then(function(response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("查询申请成功")
                    appRecords = response.data.AppRecords;
                    console.log(appRecords);
                    for (let index in appRecords) {
                        let ar = appRecords[index];
                        if (ar.ApplyStatus === 1 || ar.ReviewStatus === 2) {//已经取消的申请
                            continue;
                        }
                        let appRecordForm = createAppRecordForm();
                        appRecordForm.appRecordId = ar.AppRecordId;
                        appRecordForm.roomName = global_.Eng2ChiRoomName[ar.RoomName];
                        appRecordForm.reviewStatus = global_.ReviewStatus[ar.ReviewStatus];

                        let duration = ar.ApplyUsingTime;
                        let startTime = new Date(duration.StartTime);
                        let endTime = new Date(duration.EndTime);
                        appRecordForm.date = self_.$moment(startTime).format("YYYY-MM-DD");
                        appRecordForm.startTime = self_.$moment(startTime).format("HH:mm");
                        appRecordForm.endTime = self_.$moment(endTime).format("HH:mm");

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
        }
    },
}
</script>

<style>
.occupCard {
    width: 600px;
    margin: auto;
    padding: 20px;
}

.el-table .warning-row {
    background: oldlace;
}

.el-table .success-row {
    background: #f0f9eb;
}
</style>