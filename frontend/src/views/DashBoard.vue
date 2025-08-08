<template>
    <div class="dashboard-header">
        <h1>kuma-demo-lite</h1>
    </div>

    <div class="dashboard-action">
        <a-button type="primary" style="width: 100px" @click="getList">刷新</a-button>
    </div>

    <div class="dashboard-table">
        <a-table
            :dataSource="dataSource"
            :columns="columns"
            :pagination="false"
        />
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import axios from 'axios';

const columns = [
    { title: "#", dataIndex: "id", key: "id", align: "center", width: 50 },
    { title: "随机数", dataIndex: "random", key: "random", align: "center", width: 80 },
    { title: "创建时间", dataIndex: "created_at", key: "created_at", align: "center", width: 50 },
    { title: "更新时间", dataIndex: "updated_at", key: "updated_at", align: "center", width: 50 },
];

const dataSource = ref([]);

const getList = async () => {
    try {
        const response = await axios.get("/backend/random")
        dataSource.value = response.data
        if (dataSource.value.length > 0) {
            message.success('success');
            return;
        }
        message.error('failed');
    } catch (error) {
        console.error("获取随机数据失败:", error)
        message.error(error);
    }
};

onMounted(() => {
    getList();
});
</script>

<style scoped>
.dashboard-header {
    display: flex;
    justify-content: center;
    align-items: center;
}

.dashboard-header h1 {
    font-size: 80px;
    font-weight: 700;
    background: linear-gradient(90deg, #4a90e2, #9013fe);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    font-family: 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
    letter-spacing: 1px;
    margin: 20px 0;
    text-align: center;
}

.dashboard-action {
    display: flex;
    justify-content: center;
    margin: 30px;
}

.dashboard-table {
    display: flex;
    justify-content: center;
}

.dashboard-table > .ant-table-wrapper {
    width: 70%;
}
</style>