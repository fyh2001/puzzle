import { defineStore } from "pinia";
import { updateLogRequest } from "@/api/methods/update-log";
import type { UpdateLogReq, UpdateLogListResp } from "@/types/update-log";

export const useUpdateLogStore = defineStore("update-log", {
  state: () => ({
    logs: <UpdateLogListResp>{
      record: [],
      total: 0,
    },
  }),

  getters: {
    getLogs: (state) => {
      return state.logs.record.map((log) => {
        return {
          ...log,
          content:
            log.content instanceof Object
              ? log.content
              : JSON.parse(log.content),
          createdAt: new Date(log.createdAt).toLocaleDateString(),
          updatedAt: new Date(log.updatedAt).toLocaleDateString(),
        };
      });
    },
    getLogTotal: (state) => state.logs.total,
  },

  actions: {
    async fetchLogs(queryForm: UpdateLogReq, type: string) {
      queryForm.orderBy = "id";
      queryForm.sorted = "desc";

      const {
        data: { code, data: logsResp },
      } = await updateLogRequest.list(queryForm);

      console.log(logsResp);

      if (code === 200) {
        if (type === "append") {
          this.logs.record = [...this.logs.record, ...logsResp.record];
        }
        if (type === "replace") {
          this.logs = logsResp;
        }

        this.logs.total = logsResp.total;

        return logsResp;
      }
    },
  },
});
