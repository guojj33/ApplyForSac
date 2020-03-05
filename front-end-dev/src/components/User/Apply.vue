<template>
    <el-card :body-style="{ padding: '0px' }" class="applyCard">
      <el-form :model="appRecordForm">
          <el-form-item>
            <el-select v-model="appRecordForm.roomName" placeholder="请选择房间" class="input">
                <el-option
                    v-for="item in roomOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                    </el-option>
            </el-select>
            <el-date-picker
            class="input"
            v-model="appRecordForm.date"
            type="date"
            placeholder="请选择日期">
            </el-date-picker>
          </el-form-item>
          <el-form-item label="开始时间">
                <el-select v-model="appRecordForm.startTime" placeholder="开始时间" @change="finishSelectingStartTime()" class="input">
                    <el-option
                        v-for="item in startTimes"
                        :key="item.value"
                        :label="item.value"
                        :value="item.value">
                    </el-option>
                </el-select>
          </el-form-item>
          <el-form-item label="结束时间">
                <el-select v-model="appRecordForm.endTime" placeholder="结束时间" :disabled="canNotSelectEndTime" class="input">
                    <el-option
                        v-for="item in endTimes"
                        :key="item.value"
                        :label="item.value"
                        :value="item.value">
                    </el-option>
                </el-select>
          </el-form-item>
          <el-form-item>
                <el-input
                class="input descript"
                placeholder="申请原因"
                v-model="appRecordForm.descript"
                clearable>
                </el-input>
                <el-button type="primary" @click="CreateAppRecord()">创建</el-button>
          </el-form-item>
      </el-form>
    </el-card>
</template>

<script>
import global_ from '../../Global'
import { Loading } from 'element-ui'

export default {
    name: 'Apply',
    data() {
        return {
            roomOptions: [{
                    value: 'PianoRoom1',
                    label: '钢琴房1'
                }],
            startTimes: global_.Times.slice(0, global_.Times.length-1),
            canNotSelectEndTime: true,
            endTimes: [],
            appRecordForm: {
                roomName: '',
                date: '',
                startTime: '',
                endTime: '',
                descript: '',
            },
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
        CreateAppRecord() {
            if (this.appRecordForm.roomName === '' || this.appRecordForm.startTime === '' 
            || this.appRecordForm.endTime === '' || this.appRecordForm.descript === '' 
            || this.appRecordForm.date === '') {
                this.$alert('请补全申请表','申请');
                return;
            }
            var newAppRecordReq = {
                RoomName: '',
                Description: '',
                StartTime: '',
                EndTime: '',
            };
            newAppRecordReq.RoomName = this.appRecordForm.roomName;
            newAppRecordReq.Description = this.appRecordForm.descript;
            var dateStr = this.$moment(this.appRecordForm.date).format("YYYY-MM-DD");
            newAppRecordReq.StartTime = dateStr + "-" + this.appRecordForm.startTime + ":00";
            newAppRecordReq.EndTime = dateStr + "-" + this.appRecordForm.endTime + ":00";
            console.log(newAppRecordReq);
            var self_ = this;
            this.axios.post("/api/" + sessionStorage.getItem('AccountTypeStr') + "/" + sessionStorage.getItem('Id') + "/appRecords", newAppRecordReq)
            .then(function(response) {
                if (response.status === 200) {
                    self_.$alert('创建成功','申请',{
                        confirmButtonText: '确定',
                    });
                } else {
                    self_.$alert('创建失败','申请',{
                        confirmButtonText: '确定',
                    });
                }
                self_.appRecordForm = {
                    roomName: '',
                    date: '',
                    startTime: '',
                    endTime: '',
                    descript: '',
                }
            })
            .catch(function (error) {
                self_.$alert(error);
                self_.$alert('该时间段被占用','申请');
            })
        },
        finishSelectingStartTime() {
            this.canNotSelectEndTime = false;
            this.appRecordForm.endTime = '';
            let selectStartTimeIndex = -1;
            for (let index in global_.Times) {
                let time = global_.Times[index]
                if (time.value === this.appRecordForm.startTime) {
                    selectStartTimeIndex = index;
                    break;
                }
            }
            console.log("selectStartTimeIndex:",selectStartTimeIndex);
            let length = global_.Times.length;
            this.endTimes = global_.Times.slice(parseInt(selectStartTimeIndex) + 2, length);
        }
    }
}
</script>

<style>
.applyCard {
    width: 456px;
    padding: 20px;
    margin: auto;
}

.descript {
    width: 380px;
}

</style>