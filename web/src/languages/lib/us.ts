export default {
  tabbar: {
    practice: "Practice",
    record: "Record",
    user: "Mine",
  },
  dropdown: {
    darkMode: {
      label: "Dark Mode",
      on: "On",
      off: "Off",
    },
    theme: {
      label: "Theme Color",
      color: {
        amber: "Amber",
        green: "Green",
        indigo: "Indigo",
        rose: "Rose",
        sky: "Sky",
      },
    },
    setting: {
      label: "Settings",
    },
    updateLog: {
      label: "Update Log",
    },
    feedback: {
      label: "User Feedback",
    },
    language: {
      label: "Language",
      content: {
        cn: "中文",
        us: "English",
        jp: "日本語",
      },
    },
  },
  miniGameMap: {
    step: "Steps",
    autoPlay: "Auto Play",
    stopPlay: "Stop Play",
    last: "Last",
    next: "Next",
    reset: "Reset",
    rate: "Play Rate",
  },

  gameModeLabel: {
    practice: "Practice",
    rank: "Rank",
    battle: "Battle",
  },

  home: {
    dropdown: {
      label: "Function",
      mode: {
        label: "Game Mode",
        content: {
          practice: "Practice",
          rank: "Rank",
          battle: "Battle",
        },
      },
      dimension: {
        label: "Dimension",
      },
      colorPattern: {
        label: "Color Mode",
        content: {
          layerFirst: "Layer First",
          decreaseDimension: "Decrease Dimension",
        },
      },
    },
    step: "Steps",
    scramble: "Scramble",
  },
  record: {
    title: "Record & Rank",
    dropdown: {
      dimension: {
        label: "Dimension Selection",
      },
      mine: {
        label: "My Record",
        content: {
          practice: "Practice",
          rank: "Rank",
          battle: "Battle",
        },
      },
      rank: {
        label: "Total Rank",
        content: {
          bestSingle: "Best Single",
          bestAverage5: "Best Ao5",
          bestAverage12: "Best Ao12",
          bestStep: "Best Step",
        },
      },
      rankWeekly: {
        label: "Weekly Rank",
        content: {
          bestSingle: "Best Single",
          bestAverage5: "Best Ao5",
          bestAverage12: "Best Ao12",
          bestStep: "Best Step",
        },
      },
    },
    table: {
      column: {
        duration: "Duration",
        step: "Step",
        user: "User",
      },
    },
  },
  recordDetail: {
    title: "Record Detail",
    content: {
      gameMode: {
        label: "Game Mode",
        content: {
          practice: "Practice",
          rank: "Rank",
          battle: "Battle",
        },
      },
      type: {
        label: "Type",
        content: {
          bestSingle: "Best Single",
          bestAverage: "Average of {type}",
          bestStep: "Best Step",
        },
      },
      dimension: {
        label: "Dimension",
        content: "{dimension}x{dimension}",
      },
      rank: {
        label: "Rank",
        content: "Rank {rank}",
      },
      duration: "Duration",
      step: "Steps",
      recordBreakCount: {
        label: "Record Break Count",
        content: "{recordBreakCount} Time | {recordBreakCount} Times",
      },
      status: {
        label: "Status",
        content: {
          normal: "Normal",
          freeze: "Freezed",
          delete: "Deleted",
        },
      },
      dateTime: "Record Time",
      createdAt: "First Record",
      updatedAt: "Last Record",
      scrambleAndSolution: "Scramble & Solution",
      detailDataTitle: "No. {index} Record",
    },
  },
  mine: {
    title: "Mine",
    dropdown: {
      label: "Function",
      content: {
        editProfile: "Edit Profile",
        editPassword: "Edit Password",
        logout: {
          title: "Prompt",
          label: "Logout",
          content: "Are you sure to logout?",
          confirm: "Confirm",
          cancel: "Cancel",
          message: {
            success: "Logout Success",
            error: "Logout Failed",
          },
        },
      },
    },
    login: "Click to Login",
  },
  editProfile: {
    title: "Edit Profile",
    content: {
      avatar: "Avatar",
      username: "Username",
      nickname: "Nickname",
    },
    dialog: {
      avatar: {
        label: "Change Avatar",
        submit: "Confirm",
      },
      uploadLabel: "Select File",
    },
    messgae: {
      avatar: {
        success: "Change Avatar Success",
        error: "Change Avatar Failed",
      },
    },
  },
  login: {
    title: "Login",
    byUsername: {
      placeholder: {
        username: "Email/Phone/Username",
        password: "Please enter password",
      },
      button: {
        login: "Login",
      },
      other: {
        loginByPhone: "Login by Phone",
        register: "Register",
        forgotPassword: "Forgot Password",
      },
      rules: {
        rulesError: "Please fill in according to the rules",
        username: {
          label: "Username",
          message: {
            required: "Please enter username",
            pattern: "Username only allows 6-12 digits or letters",
          },
        },
        password: {
          label: "Password",
          message: {
            required: "Please enter password",
            pattern: "Password only allows 6-12 digits or letters",
          },
        },
      },
    },
    message: {
      success: "Login Success",
      error: "Login Failed",
    },
  },
  register: {
    title: "Register",
    byUsername: {
      placeholder: {
        username: "Username for login",
        nickname: "Nickname displayed to users",
        password: "Password for login",
      },
      button: {
        register: "Register",
      },
      other: {
        loginByPhone: "Login by phone",
        login: "Login now",
        forgotPassword: "Forgot password",
      },
      rules: {
        rulesError: "Please follow the rules",
        username: {
          label: "Username",
          message: {
            required: "Please enter a username",
            pattern:
              "Username can only be a combination of 6-12 letters or numbers",
          },
        },
        password: {
          label: "Password",
          message: {
            required: "Please enter a password",
            pattern:
              "Password can only be a combination of 6-12 letters or numbers",
          },
        },
        nickname: {
          label: "Nickname",
          message: {
            required: "Please enter a nickname",
            pattern:
              "Nickname cannot contain characters other than {'@'} # ,.?，。？ and must be 2-15 characters long",
          },
        },
      },
    },
    message: {
      success: "Registration successful",
      error: "Registration failed",
    },
  },
};
