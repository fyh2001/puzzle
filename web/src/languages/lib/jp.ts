export default {
  tabbar: {
    practice: "練習",
    record: "記録",
    user: "私の",
  },
  dropdown: {
    notification: {
      label: "通知",
    },
    darkMode: {
      label: "ダークモード",
      on: "オン",
      off: "オフ",
    },
    theme: {
      label: "テーマ色",
      color: {
        amber: "琥珀",
        green: "翠緑",
        indigo: "藍色",
        rose: "ローズ",
        sky: "空",
      },
    },
    setting: {
      label: "設定",
    },
    updateLog: {
      label: "更新ログ",
    },
    feedback: {
      label: "フィードバック",
    },
    language: {
      label: "言語",
      content: {
        cn: "中文",
        us: "English",
        jp: "日本語",
      },
    },
  },
  miniGameMap: {
    step: "ステップ",
    autoPlay: "自動再生",
    stopPlay: "一時停止",
    last: "前へ",
    next: "次へ",
    reset: "リセット",
    rate: "再生速度",
  },

  gameModeLabel: {
    practice: "練習",
    rank: "ランキング",
    battle: "バトル",
  },

  home: {
    dropdown: {
      label: "機能",
      mode: {
        label: "ゲームモード",
        content: {
          practice: "練習",
          rank: "ランキング",
          battle: "バトル",
        },
      },
      dimension: {
        label: "次元選択",
      },
      colorPattern: {
        label: "カラーパターン",
        content: {
          layerFirst: "レイヤー優先",
          decreaseDimension: "次元を下げる",
        },
      },
    },
    step: "ステップ",
    scramble: "スクランブル",
  },
  record: {
    title: "記録とランキング",
    dropdown: {
      dimension: {
        label: "次元選択",
      },
      mine: {
        label: "私の記録",
        content: {
          practice: "練習",
          rank: "ランキング",
          battle: "バトル",
        },
      },
      rank: {
        label: "全体ランキング",
        content: {
          bestSingle: "最高のシングル",
          bestAverage5: "最高の5回平均",
          bestAverage12: "最高の12回平均",
          bestStep: "最高のステップ",
        },
      },
      rankWeekly: {
        label: "週間ランキング",
        content: {
          bestSingle: "最高のシングル",
          bestAverage5: "最高の5回平均",
          bestAverage12: "最高の12回平均",
          bestStep: "最高のステップ",
        },
      },
    },
    table: {
      column: {
        duration: "期間",
        step: "ステップ",
        user: "ユーザー",
      },
    },
  },
  recordDetail: {
    title: "記録詳細",
    content: {
      gameMode: {
        label: "ゲームモード",
        content: {
          practice: "練習",
          rank: "ランキング",
          battle: "バトル",
        },
      },
      type: {
        label: "タイプ",
        content: {
          bestSingle: "最高のシングル",
          bestAverage: "最高の {type} 回平均",
          bestStep: "最高のステップ",
        },
      },
      dimension: {
        label: "次元",
        content: "{dimension} 次元",
      },
      rank: {
        label: "ランキング",
        content: "第 {rank} 位",
      },
      duration: "期間",
      step: "ステップ",
      recordBreakCount: {
        label: "記録更新回数",
        content: "{recordBreakCount} 回",
      },
      status: {
        label: "ステータス",
        content: {
          normal: "正常",
          freeze: "フリーズ",
          delete: "削除",
        },
      },
      dateTime: "記録時間",
      createdAt: "初回記録",
      updatedAt: "最終記録",
      scrambleAndSolution: "スクランブルと解決",
      detailDataTitle: "第 {index} 回の記録",
    },
  },
  mine: {
    title: "私の",
    dropdown: {
      label: "機能",
      content: {
        editProfile: "プロフィールを編集",
        editPassword: "パスワードを変更",
        logout: {
          title: "ヒント",
          label: "ログアウト",
          content: "本当にログアウトしますか？",
          confirm: "確認",
          cancel: "キャンセル",
          message: {
            success: "ログアウト成功",
            error: "ログアウト失敗",
          },
        },
      },
    },
    login: "ログインする",
  },
  editProfile: {
    title: "プロフィールを編集",
    content: {
      avatar: "アバター",
      username: "ユーザー名",
      nickname: "ニックネーム",
    },
    dialog: {
      avatar: {
        label: "アバターを変更",
        button: {
          confirm: "確認",
        },
      },
      username: {
        label: "ユーザー名を変更",
        input: {
          placeholder: "ユーザー名を入力してください",
        },
        button: {
          confirm: "確認",
        },
      },
      nickname: {
        label: "ニックネームを変更",
        input: {
          placeholder: "ニックネームを入力してください",
        },
        button: {
          confirm: "確認",
          cancel: "キャンセル",
        },
      },
      uploadLabel: "ファイルを選択",
    },
    messgae: {
      success: "編集成功",
      error: "編集失敗",
    },
  },
  login: {
    title: "ログイン",
    byUsername: {
      placeholder: {
        username: "メール/携帯電話番号/ユーザー名",
        password: "パスワードを入力してください",
      },
      button: {
        login: "ログイン",
      },
      other: {
        loginByPhone: "携帯電話でログイン",
        register: "すぐに登録",
        forgotPassword: "パスワードを忘れた",
      },
      rules: {
        rulesError: "ルールに従って記入してください",
        username: {
          label: "ユーザー名",
          message: {
            required: "ユーザー名を入力してください",
            pattern: "ユーザー名は6-12文字の英数字のみが許可されます",
          },
        },
        password: {
          label: "パスワード",
          message: {
            required: "パスワードを入力してください",
            pattern: "パスワードは6-12文字の英数字のみが許可されます",
          },
        },
      },
    },
    message: {
      success: "ログイン成功",
      error: "ログイン失敗",
    },
  },
  register: {
    title: "登録",
    byUsername: {
      placeholder: {
        username: "ログイン用のユーザー名",
        nickname: "ユーザーに表示されるニックネーム",
        password: "ログイン用のパスワード",
      },
      button: {
        register: "登録",
      },
      other: {
        loginByPhone: "携帯電話でログイン",
        login: "すぐにログイン",
        forgotPassword: "パスワードを忘れた",
      },
      rules: {
        rulesError: "ルールに従って記入してください",
        username: {
          label: "ユーザー名",
          message: {
            required: "ユーザー名を入力してください",
            pattern: "ユーザー名は6-12文字の英数字のみが許可されます",
          },
        },
        password: {
          label: "パスワード",
          message: {
            required: "パスワードを入力してください",
            pattern: "パスワードは6-12文字の英数字のみが許可されます",
          },
        },
        nickname: {
          label: "ニックネーム",
          message: {
            required: "ニックネームを入力してください",
            pattern:
              "{'@'}#,.?,. 以外のニックネームは使用できません。 ? 以外の文字で、長さは2～10桁（漢字の場合は6桁）",
          },
        },
      },
    },
    message: {
      success: "登録成功",
      error: "登録失敗",
    },
  },
};
