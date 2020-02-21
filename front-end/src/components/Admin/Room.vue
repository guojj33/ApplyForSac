<template>
    <el-card :body-style="{ padding: '0px' }" class="roomCard">
        <el-form :model="createForm">
            <el-form-item>
                <el-input
                placeholder="新房间名"
                v-model="createForm.roomName"
                clearable
                class="inputBox">
                </el-input>
                <el-button
                type="primary"
                @click="CreateRoom()"
                size="large">
                    创建房间
                </el-button>
            </el-form-item>
        </el-form>
        <el-table
        v-loading="isLoading"
        :data="rooms"
        stype="width: 100%">
            <el-table-column
            prop="roomName"
            label="房间名"
            width="200px">
            </el-table-column>
        </el-table>   
    </el-card>
</template>>

<script>
export default {
    name: 'room',
    data() {
        return {
            isLoading: true,
            rooms: [{
                roomName: '',
            }],
            createForm: {
                roomName: '',
            }
        }
    },
    created(){
        this.GetRoomOptions();
    },
    methods: {
        GetRoomOptions() {
            var createRoomOption = function() {
                var roomOption = new Object ({
                    roomName: '',
                });
                return roomOption
            };
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
                        roomOption.roomName = r.RoomName
                        options.push(roomOption);
                    }
                    console.log(options);
                    self_.rooms = options
                } else {
                    console.log("获取房间表失败");
                }
                self_.isLoading = false;
            })
            .catch(function(error) {
                alert(error);
            })
        },
        CreateRoom() {
            if (this.createForm.roomName === '') {
                this.$alert('房间名不能为空','房间管理');
                return;
            }
            let self_ = this;
            this.axios.post("/api/rooms", this.createForm)
            .then(function(response) {
                if (response.status === 200) {
                    self_.$alert('创建成功','房间管理');
                    self_.GetRoomOptions();
                } else {
                    self_.$alert('创建失败','房间管理');
                }
                self_.createForm = {
                    roomName: '',
                }
            })
            .catch(function(error) {
                alert(error);
            })
        }
    }
}
</script>

<style>
.roomCard {
    width: 250px;
    padding: 20px;
    margin: auto;
}
.inputBox {
    width: 140px;
}
</style>