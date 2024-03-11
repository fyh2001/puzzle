import { Pagination } from "@/types/pagination";

// 通知
export interface Notification {
  id: string;
  userId: string;
  type: number;
  content: string;
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
  status?: number | null;
  dateRange?: [string, string] | [Date, Date];
  pagination?: Pagination;
  sorted?: string;
  orderBy?: string;
}

export interface NotificationResp {
  id: string;
  userId: string;
  type: number;
  content: string;
  status: number;
  notificationTypeInfo: NotificationTypeResp;
  notificationUserStatusInfo: NotificationUserStatusResp;
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

// 用户通知状态
export interface NotificationUserStatus {
  id: string;
  userId: string;
  notificationId: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface NotificationUserStatusReq {
  id?: string;
  ids?: string[];
  userId?: string;
  username?: string;
  nickname?: string;
  notificationId?: string;
  status?: number | null;
  dateRange?: [string, string] | [Date, Date];
  pagination?: Pagination;
  sorted?: string;
  orderBy?: string;
}

export interface NotificationUserStatusResp {
  id: string;
  userId: string;
  notificationId: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface NotificationUserStatusListResp {
  total: number;
  records: NotificationUserStatusResp[];
}
