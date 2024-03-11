import { defineStore } from "pinia";
import {
  authRequest,
  userManageRequest,
  recordManageRequest,
  recordBestSingleRequest,
  recordBestAverageRequest,
  recordBestStepRequest,
} from "api/admin";
import { formatDurationInRecord } from "@/utils/time";
import type { AuthReq, AdminData } from "@/types/admin";
import type { UserReq, UserResp } from "@/types/user";
import type {
  RecordResp,
  RecordReq,
  RecordBestSingleReq,
  RecordBestSingleResp,
  RecordBestAverageReq,
  RecordBestAverageResp,
  RecordBestStepReq,
  RecordBestStepResp,
} from "@/types/record";

export const useAdminStore = defineStore("admin", {
  persist: true,

  state: () => ({
    authorization: {
      token: "",
      secretUrl: "",
      expiredAt: 0,
    },
    data: {
      user: <AdminData<UserResp>>{
        list: [],
        total: 0,
      },
      personRecord: <AdminData<RecordResp>>{
        list: [],
        total: 0,
      },
      bestSingleRecord: <AdminData<RecordBestSingleResp>>{
        list: [],
        total: 0,
      },
      bestAverageRecord: <AdminData<RecordBestAverageResp>>{
        list: [],
        total: 0,
      },
      bestStepRecord: <AdminData<RecordBestStepResp>>{
        list: [],
        total: 0,
      },
    },
  }),

  getters: {
    getToken: (state) => {
      if (state.authorization.expiredAt < Date.now()) {
        return "";
      }

      return state.authorization.token;
    },
    getSecretUrl: (state) => state.authorization.secretUrl,
    getUserList: (state) => state.data.user.list,
    getUserTotal: (state) => state.data.user.total,
    getPersonRecordList: (state) => {
      return state.data.personRecord.list.map((record) => {
        const durationFormat = formatDurationInRecord(record.duration);

        return {
          ...record,
          durationFormat,
        };
      });
    },
    getPersonRecordTotal: (state) => state.data.personRecord.total,
    getBestSingleRecordList: (state) => {
      return state.data.bestSingleRecord.list.map((record) => {
        const durationFormat = formatDurationInRecord(record.recordDuration);
        const recordDetail = record.recordDetail.records.map((record) => {
          const durationFormat = formatDurationInRecord(record.duration);
          return {
            ...record,
            durationFormat,
          };
        });

        return {
          ...record,
          durationFormat,
          recordDetail,
        };
      });
    },
    getBestSingleRecordTotal: (state) => state.data.bestSingleRecord.total,
    getBestAverageRecordList: (state) => {
      return state.data.bestAverageRecord.list.map((record) => {
        const durationFormat = formatDurationInRecord(
          record.recordAverageDuration
        );
        const recordDetail = record.recordDetail.records.map((record) => {
          const durationFormat = formatDurationInRecord(record.duration);
          return {
            ...record,
            durationFormat,
          };
        });

        return {
          ...record,
          durationFormat,
          recordDetail,
        };
      });
    },
    getBestAverageRecordTotal: (state) => state.data.bestAverageRecord.total,
    getBestStepRecordList: (state) => {
      return state.data.bestStepRecord.list.map((record) => {
        const recordDetail = record.recordDetail.records.map((record) => {
          const durationFormat = formatDurationInRecord(record.duration);
          return {
            ...record,
            durationFormat,
          };
        });

        return {
          ...record,
          recordDetail,
        };
      });
    },
    getBestStepRecordTotal: (state) => state.data.bestStepRecord.total,
  },

  actions: {
    // 登录
    async auth(data: AuthReq): Promise<boolean> {
      const {
        data: { code, data: authResp },
      } = await authRequest.authorization(data);

      if (code === 200) {
        this.authorization.token = authResp.token;
        this.authorization.secretUrl = authResp.secretUrl;
        this.authorization.expiredAt = Date.now() + 24 * 60 * 60 * 1000; // 24小时
        return true;
      }

      return false;
    },

    // 获取二维码
    async fetchSecretUrl() {
      const {
        data: { code, data: secretUrlResp },
      } = await authRequest.getUrl();

      if (code === 200) {
        this.authorization.secretUrl = secretUrlResp.secretUrl;
      }
    },

    // 重置二维码
    async resetOtp() {
      const {
        data: { code, data: secretUrlResp },
      } = await authRequest.resetOtp();

      if (code === 200) {
        this.authorization.secretUrl = secretUrlResp.secretUrl;
      }
    },

    // 重置信息
    resetaAuthorization() {
      this.authorization.token = "";
      this.authorization.secretUrl = "";
      this.authorization.expiredAt = 0;
    },

    // 获取用户列表
    async fetchUserList(queryForm: UserReq) {
      const {
        data: { code, data: userListResp, msg },
      } = await userManageRequest.list(queryForm);

      if (code === 200) {
        this.data.user.list = userListResp.records;
        this.data.user.total = userListResp.total;
      }

      return { code, msg };
    },
    // 更新用户列表
    async updateUserList(data: UserReq) {
      const {
        data: { code, msg, data: updateResp },
      } = await userManageRequest.update(data);

      return { code, msg, data: updateResp };
    },

    // 获取个人记录列表
    async fetchRecordList(queryForm: RecordReq) {
      const {
        data: { code, data: recordListResp, msg },
      } = await recordManageRequest.list(queryForm);

      if (code === 200) {
        console.log("recordList", recordListResp);

        this.data.personRecord.list = recordListResp.records;
        this.data.personRecord.total = recordListResp.total;
      }

      return { code, msg };
    },
    // 更新记录列表
    async updateRecordList(data: RecordReq) {
      const {
        data: { code, msg, data: updateResp },
      } = await recordManageRequest.update(data);

      return { code, msg, data: updateResp };
    },

    // 获取最佳单次记录列表
    async fetchRecordBestSingleList(queryForm: RecordBestSingleReq) {
      const {
        data: { code, data: recordListResp, msg },
      } = await recordBestSingleRequest.list(queryForm);

      console.log(recordListResp);

      if (code === 200) {
        this.data.bestSingleRecord.list = recordListResp.records;
        this.data.bestSingleRecord.total = recordListResp.total;
      }

      return { code, msg };
    },

    // 获取最佳平均记录列表
    async fetchRecordBestAverageList(queryForm: RecordBestAverageReq) {
      const {
        data: { code, data: recordListResp, msg },
      } = await recordBestAverageRequest.list(queryForm);

      if (code === 200) {
        this.data.bestAverageRecord.list = recordListResp.records;
        this.data.bestAverageRecord.total = recordListResp.total;
      }
      return { code, msg };
    },

    // 获取最佳步数记录列表
    async fetchRecordBestStepList(queryForm: RecordBestStepReq) {
      const {
        data: { code, data: recordListResp, msg },
      } = await recordBestStepRequest.list(queryForm);

      if (code === 200) {
        this.data.bestStepRecord.list = recordListResp.records;
        this.data.bestStepRecord.total = recordListResp.total;
      }

      return { code, msg };
    },
  },
});
