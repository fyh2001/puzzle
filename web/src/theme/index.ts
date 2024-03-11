const common = {
  Message: {
    borderRadius: "0.5rem",
  },
  Dialog: {
    borderRadius: "0.75rem",
  },
};

export const themeProvider = (other: any) => {
  return {
    ...common,
    ...other,
  };
};
