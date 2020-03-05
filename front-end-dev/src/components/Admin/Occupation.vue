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
      prop="userType"
      label="用户类型"
      width="100%">
      </el-table-column>
      <el-table-column
      prop="descript"
      label="原因"
      width="100%">
      </el-table-column>
      <el-table-column
      prop="date"
      column-key="date"
      sortable
      :filter-method="queryByDate"
      :filters="[{text:'筛选',value:'query'}]"
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
import { Loading } from 'element-ui'

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
    created(){
        this.GetRoomOptions();
    },
    methods: {
        GetRoomOptions() {
            var createRoomOption = function() {
                var roomOption = new Object ({
                    value: '',
                    label: '',
                });
                return roomOption
            };
            let loadingInst = Loading.service();
            let self_ = this;
            let options = [];
            this.axios.get("/api/rooms")
            .then(function(response) {
                console.log(response);
                if (response.status === 200) {
                    console.log("获取房间表成功");
                    let rooms = response.data.rooms;
                    for (let index in rooms) {
                        let r = rooms[index];
                        console.log(r);
                        let roomOption = createRoomOption();
                        roomOption.value = r.RoomName;
                        roomOption.label = r.roomName;
                        options.push(roomOption);
                    }
                    console.log(options);
                    self_.roomOptions = options
                } else {
                    console.log("获取房间表失败");
                }
                loadingInst.close();
            })
            .catch(function(error) {
                alert(error);
            })
        },
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
                    userType: '',
                    descript: '',
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
                        appRecordForm.roomName = ar.RoomName;
                        appRecordForm.reviewStatus = global_.ReviewStatus[ar.ReviewStatus];

                        let duration = ar.ApplyUsingTime;
                        let startTime = new Date(duration.StartTime);
                        let endTime = new Date(duration.EndTime);
                        appRecordForm.date = self_.$moment(startTime).utc().format("YYYY-MM-DD");
                        appRecordForm.startTime = self_.$moment(startTime).utc().format("HH:mm");
                        appRecordForm.endTime = self_.$moment(endTime).utc().format("HH:mm");
                        if (ar.ApplyUserId === "SAC") {
                            appRecordForm.userType = "管理员";
                        } else {
                            appRecordForm.userType = "普通用户";
                        }
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
        queryByDate(value, row, column) {
            if (this.searchForm.selectedDate === '') {
                return true;
            }
            let property = column['property'];
            let date = row[property];
            let selectedDateStr = this.$moment(this.searchForm.selectedDate).format("YYYY-MM-DD");
            return date === selectedDateStr;
        }
    },
}
</script>

<style>
.occupCard {
    width: 810px;
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