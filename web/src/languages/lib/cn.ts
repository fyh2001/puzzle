export default {
  tabbar: {
    practice: "练习",
    record: "记录",
    user: "我的",
  },
  dropdown: {
    notification: {
      label: "通知",
    },
    darkMode: {
      label: "深色模式",
      on: "开启",
      off: "关闭",
    },
    theme: {
      label: "主题色",
      color: {
        amber: "琥珀",
        green: "翠绿",
        indigo: "靛蓝",
        rose: "玫瑰",
        sky: "天空",
      },
    },
    setting: {
      label: "设置",
    },
    updateLog: {
      label: "更新日志",
    },
    feedback: {
      label: "用户反馈",
    },
    language: {
      label: "语言",
      content: {
        cn: "中文",
        us: "English",
        jp: "日本語",
      },
    },
  },
  miniGameMap: {
    step: "步数",
    autoPlay: "自动播放",
    stopPlay: "暂停播放",
    last: "上一步",
    next: "下一步",
    reset: "复 原",
    rate: "播放速度",
  },

  gameModeLabel: {
    practice: "练习",
    rank: "排位",
    battle: "对战",
  },

  home: {
    dropdown: {
      label: "功能",
      mode: {
        label: "游戏模式",
        content: {
          practice: "练习",
          rank: "排位",
          battle: "对战",
        },
      },
      dimension: {
        label: "阶数选择",
      },
      colorPattern: {
        label: "颜色模式",
        content: {
          layerFirst: "层先",
          decreaseDimension: "降阶",
        },
      },
    },
    step: "步数",
    scramble: "打乱",
  },
  record: {
    title: "记录与排名",
    dropdown: {
      dimension: {
        label: "阶数选择",
      },
      mine: {
        label: "我的记录",
        content: {
          practice: "练习",
          rank: "排名",
          battle: "对战",
        },
      },
      rank: {
        label: "总排名",
        content: {
          bestSingle: "最佳单次",
          bestAverage5: "最佳5次平均",
          bestAverage12: "最佳12次平均",
          bestStep: "最佳步数",
        },
      },
      rankWeekly: {
        label: "周排名",
        content: {
          bestSingle: "最佳单次",
          bestAverage5: "最佳5次平均",
          bestAverage12: "最佳12次平均",
          bestStep: "最佳步数",
        },
      },
    },
    table: {
      column: {
        duration: "时长",
        step: "步数",
        user: "用户",
      },
    },
  },
  recordDetail: {
    title: "记录详情",
    content: {
      gameMode: {
        label: "游戏模式",
        content: {
          practice: "练习",
          rank: "排位",
          battle: "对战",
        },
      },
      type: {
        label: "类型",
        content: {
          bestSingle: "最佳单次",
          bestAverage: "最佳 {type} 次平均",
          bestStep: "最佳步数",
        },
      },
      dimension: {
        label: "阶数",
        content: "{dimension} 阶",
      },
      rank: {
        label: "排名",
        content: "第 {rank} 名",
      },
      duration: "耗时",
      step: "步数",
      recordBreakCount: {
        label: "记录破纪录次数",
        content: "{recordBreakCount} 次",
      },
      status: {
        label: "状态",
        content: {
          normal: "正常",
          freeze: "冻结",
          delete: "删除",
        },
      },
      dateTime: "记录时间",
      createdAt: "初次记录",
      updatedAt: "上次记录",
      scrambleAndSolution: "打乱与还原",
      detailDataTitle: "第 {index} 次记录",
    },
  },
  mine: {
    title: "我的",
    dropdown: {
      label: "功能",
      content: {
        editProfile: "编辑资料",
        editPassword: "修改密码",
        logout: {
          title: "提示",
          label: "退出登录",
          content: "确认退出登录吗？",
          confirm: "确认",
          cancel: "取消",
          message: {
            success: "退出登录成功",
            error: "退出登录失败",
          },
        },
      },
    },
    login: "点我登录",
  },
  editProfile: {
    title: "编辑资料",
    content: {
      avatar: "头像",
      username: "用户名",
      nickname: "昵称",
    },
    dialog: {
      avatar: {
        label: "修改头像",
        button: {
          confirm: "确认修改",
        },
      },
      username: {
        label: "修改用户名",
        input: {
          placeholder: "请输入用户名",
        },
        button: {
          confirm: "确认",
        },
      },
      nickname: {
        label: "修改昵称",
        input: {
          placeholder: "请输入昵称",
        },
        button: {
          confirm: "确认",
          cancel: "取消",
        },
      },
      uploadLabel: "选择文件",
    },

    messgae: {
      success: "修改成功",
      error: "修改失败",
    },
  },
  login: {
    title: "登录",
    byUsername: {
      placeholder: {
        username: "邮箱/手机号码/用户名",
        password: "请输入密码",
      },
      button: {
        login: "登录",
      },
      other: {
        loginByPhone: "手机登录",
        register: "立即注册",
        forgotPassword: "忘记密码",
      },
      rules: {
        rulesError: "请按照规则填写",
        username: {
          label: "用户名",
          message: {
            required: "请输入用户名",
            pattern: "用户名只允许是6-12位字母或数字的组合",
          },
        },
        password: {
          label: "密码",
          message: {
            required: "请输入密码",
            pattern: "密码只允许是6-12位字母或数字的组合",
          },
        },
      },
    },
    message: {
      success: "登录成功",
      error: "登录失败",
    },
  },
  register: {
    title: "注册",
    byUsername: {
      placeholder: {
        username: "用于登录的用户名",
        nickname: "用户展示的昵称",
        password: "用于登录的密码",
      },
      button: {
        register: "注册",
      },
      other: {
        loginByPhone: "手机登录",
        login: "立即登录",
        forgotPassword: "忘记密码",
      },
      rules: {
        rulesError: "请按照规则填写",
        username: {
          label: "用户名",
          message: {
            required: "请输入用户名",
            pattern: "用户名只允许是6-12位字母或数字的组合",
          },
        },
        password: {
          label: "密码",
          message: {
            required: "请输入密码",
            pattern: "密码只允许是6-12位字母或数字的组合",
          },
        },
        nickname: {
          label: "昵称",
          message: {
            required: "请输入昵称",
            pattern:
              "昵称不允许出现除了{'@'}# ,.?，。？以外的字符，且长度为2-10位(中文6位)",
          },
        },
      },
    },
    message: {
      success: "注册成功",
      error: "注册失败",
    },
  },
};
