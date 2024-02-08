import { defineStore } from "pinia";
import { recordRequest } from "api/record";
import { getAverageOf5, getAverageOf12 } from "@/utils/record";
import { formatDurationInRecord } from "@/utils/time";
import type { RecordListResp, RecordReq } from "@/types/record";

export const useRecordStore = defineStore("record", {
  state: () => ({
    // 个人记录
    personRecord: <RecordListResp>{ records: [], total: 0 },
    // 排行榜
    rankedRecord: {
      bestSingle: <RecordListResp>{ records: [], total: 0 }, //最佳单次
      bestAverage5: <RecordListResp>{ records: [], total: 0 }, //最佳5次
      bestAverage12: <RecordListResp>{ records: [], total: 0 }, //最佳12次
      bestStep: <RecordListResp>{ records: [], total: 0 }, //最佳步数
    },
    // 周排行榜
    weeklyRankedRecord: {
      bestSingle: <RecordListResp>{ records: [], total: 0 }, //最佳单次
      bestAverage5: <RecordListResp>{ records: [], total: 0 }, //最佳5次
      bestAverage12: <RecordListResp>{ records: [], total: 0 }, //最佳12次
      bestStep: <RecordListResp>{ records: [], total: 0 }, //最佳步数
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
  },

  actions: {
    // 获取个人记录
    async getPersonRecords(queryForm: RecordReq) {
      const {
        data: { code, data: recordList },
      } = await recordRequest.list({
        ...queryForm,
        type: 0,
        sorted: "desc",
      });

      if (code === 200) {
        this.personRecord = recordList;
        console.log("个人记录", recordList);
      }
    },
    // async getBestSingle(queryForm: RecordReq) {},
  },
});
