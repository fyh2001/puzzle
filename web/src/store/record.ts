import { defineStore } from "pinia";
import { recordRequest } from "api/record";
import { getAverageOf5, getAverageOf12 } from "@/utils/record";
import { formatDurationInRecord } from "@/utils/time";
import type {
  RecordReq,
  RecordListResp,
  RecordBestSingleReq,
  RecordBestSingleListResp,
  RecordBestAverageReq,
  RecordBestAverageListResp,
  RecordBestStepReq,
  RecordBestStepListResp,
} from "@/types/record";

export const useRecordStore = defineStore("record", {
  state: () => ({
    // 个人记录
    personRecord: <RecordListResp>{ records: [], total: 0 },
    // 排行榜
    rankedRecord: {
      bestSingle: <RecordBestSingleListResp>{ records: [], total: 0 }, //最佳单次
      bestAverage5: <RecordBestAverageListResp>{ records: [], total: 0 }, //最佳5次
      bestAverage12: <RecordBestAverageListResp>{ records: [], total: 0 }, //最佳12次
      bestStep: <RecordBestStepListResp>{ records: [], total: 0 }, //最佳步数
    },
    // 周排行榜
    weeklyRankedRecord: {
      bestSingle: <RecordBestSingleListResp>{ records: [], total: 0 }, //最佳单次
      bestAverage5: <RecordBestAverageListResp>{ records: [], total: 0 }, //最佳5次
      bestAverage12: <RecordBestAverageListResp>{ records: [], total: 0 }, //最佳12次
      bestStep: <RecordBestStepListResp>{ records: [], total: 0 }, //最佳步数
    },
  }),

  getters: {
    getPersonRecordList: (state) => {
      return state.personRecord.records
        .map((record, recordIndex) => {
          const ao5 = getAverageOf5(
            state.personRecord.records.slice(recordIndex, recordIndex + 5)
          );
          const ao12 = getAverageOf12(
            state.personRecord.records.slice(recordIndex, recordIndex + 12)
          );
          const duration = formatDurationInRecord(record.duration);

          return {
            ...record,
            ao5,
            ao12,
            duration,
          };
        })
        .slice(0, 10);
    },
    getBestSingleRecordList: (state) => {
      return state.rankedRecord.bestSingle.records.map((record) => {
        const duration = formatDurationInRecord(record.recordDuration);
        const step = record.recordStep;
        const nickname = record.userInfo?.nickname;
        return {
          ...record,
          duration,
          step,
          nickname,
        };
      });
    },
    getBestAverage5RecordList: (state) => {
      return state.rankedRecord.bestAverage5.records.map((record) => {
        const duration = formatDurationInRecord(record.recordAverageDuration);
        const nickname = record.userInfo?.nickname;
        return {
          ...record,
          duration,
          nickname,
        };
      });
    },
    getBestAverage12RecordList: (state) => {
      return state.rankedRecord.bestAverage12.records.map((record) => {
        const duration = formatDurationInRecord(record.recordAverageDuration);
        const nickname = record.userInfo?.nickname;
        return {
          ...record,
          duration,
          nickname,
        };
      });
    },
    getBestStepRecordList: (state) => {
      return state.rankedRecord.bestStep.records.map((record) => {
        const step = record.recordStep;
        const nickname = record.userInfo?.nickname;
        return {
          ...record,
          step,
          nickname,
        };
      });
    },
  },

  actions: {
    // 获取个人记录
    async getPersonRecords(queryForm: RecordReq) {
      const {
        data: { code, data: recordList },
      } = await recordRequest.listRecord({
        ...queryForm,
        sorted: "desc",
      });

      if (code === 200) {
        this.personRecord = recordList;
      }
    },
    // 获取最佳单次记录
    async getBestSingleRecords(queryForm: RecordBestSingleReq) {
      const {
        data: { code, data: recordList },
      } = await recordRequest.listBestSingle(queryForm);

      if (code === 200) {
        this.rankedRecord.bestSingle = recordList;
      }
    },

    // 获取最佳平均5次记录
    async getBestAverage5Records(queryForm: RecordBestAverageReq) {
      const {
        data: { code, data: recordList },
      } = await recordRequest.listBestAverage({
        ...queryForm,
        type: 5,
      });

      if (code === 200) {
        console.log(recordList);
        this.rankedRecord.bestAverage5 = recordList;
      }
    },

    // 获取最佳平均12次记录
    async getBestAverage12Records(queryForm: RecordBestAverageReq) {
      const {
        data: { code, data: recordList },
      } = await recordRequest.listBestAverage({
        ...queryForm,
        type: 12,
      });

      if (code === 200) {
        this.rankedRecord.bestAverage12 = recordList;
      }
    },

    // 获取最佳步数记录
    async getBestStepRecords(queryForm: RecordBestStepReq) {
      const {
        data: { code, data: recordList },
      } = await recordRequest.listBestStep(queryForm);

      if (code === 200) {
        this.rankedRecord.bestStep = recordList;
      }
    },
  },
});
