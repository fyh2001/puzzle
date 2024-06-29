import { Pagination } from "@/types/pagination";

// 通知
export interface Notification {
  id: string;
  userId: string;
  type: number;
  content: string;
  readStatus: number;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface NotificationReq {
  id?: string;
  ids?: string[];
  userId?: string;
  username?: string;
  nickname?: string;
  type?: number | null;
  content?: string;
  readStatus?: number | null;
  status?: number | null;
  dateRange?: [string, string] | [Date, Date];
  pagination?: Pagination;
  sorted?: string;
  orderBy?: string;
  onlyTotal?: boolean;
}

export interface NotificationResp {
  id: string;
  userId: string;
  type: number;
  content: string;
  readStatus: number;
  status: number;
  notificationTypeInfo: NotificationTypeResp;
  createdAt: string;
  updatedAt: string;
}

export interface NotificationListResp {
  total: number;
  records: NotificationResp[];
}

// 通知类型
export interface NotificationType {
  id: number;
  name: string;
  icon: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface NotificationTypeReq {
  id?: string;
  ids?: string[];
  name?: string;
  icon?: string;
  status?: number | null;
  dateRange?: [string, string] | [Date, Date];
  pagination?: Pagination;
  sorted?: string;
  orderBy?: string;
}

export interface NotificationTypeResp {
  id: number;
  name: string;
  icon: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface NotificationTypeListResp {
  total: number;
  records: NotificationTypeResp[];
}
